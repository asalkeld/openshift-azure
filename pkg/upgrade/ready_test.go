package upgrade

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestIsPodReady(t *testing.T) {
	tests := []struct {
		name string
		pod  *corev1.Pod
		want bool
	}{
		{
			name: "ready",
			want: true,
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name: "missing",
			want: false,
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		if got := isPodReady(tt.pod); got != tt.want {
			t.Errorf("isPodReady(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNodeReady(t *testing.T) {
	tests := []struct {
		name string
		node *corev1.Node
		want bool
	}{
		{
			name: "ready",
			want: true,
			node: &corev1.Node{
				Status: corev1.NodeStatus{
					Conditions: []corev1.NodeCondition{
						{
							Type:   corev1.NodeOutOfDisk,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.NodeReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name: "missing",
			want: false,
			node: &corev1.Node{
				Status: corev1.NodeStatus{
					Conditions: []corev1.NodeCondition{
						{
							Type:   corev1.NodeOutOfDisk,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		if got := isNodeReady(tt.node); got != tt.want {
			t.Errorf("isNodeReady(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNodeIsReady(t *testing.T) {
	tests := []struct {
		name     string
		kc       kubernetes.Interface
		nodeName string
		want     bool
		wantErr  bool
	}{
		{
			name:     "not found",
			nodeName: "toad",
			wantErr:  false,
			want:     false,
			kc:       fake.NewSimpleClientset(),
		},
		{
			name:     "not ready",
			nodeName: "toad",
			wantErr:  false,
			want:     false,
			kc: fake.NewSimpleClientset(&corev1.Node{
				TypeMeta: metav1.TypeMeta{
					Kind: "node",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "toad",
				},
			}),
		},
		{
			name:     "ready",
			nodeName: "toad",
			wantErr:  false,
			want:     true,
			kc: fake.NewSimpleClientset(&corev1.Node{
				TypeMeta: metav1.TypeMeta{
					Kind: "node",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "toad",
				},
				Status: corev1.NodeStatus{
					Conditions: []corev1.NodeCondition{
						{
							Type:   corev1.NodeReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}),
		},
	}
	for _, tt := range tests {
		got, err := nodeIsReady(tt.kc, tt.nodeName)
		if (err != nil) != tt.wantErr {
			t.Errorf("nodeIsReady() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if got != tt.want {
			t.Errorf("nodeIsReady() = %v, want %v", got, tt.want)
		}
	}
}

func TestMasterIsReady(t *testing.T) {
	tests := []struct {
		name     string
		kc       kubernetes.Interface
		nodeName string
		want     bool
		wantErr  bool
	}{
		{
			name:     "node not found",
			nodeName: "toad",
			wantErr:  false,
			want:     false,
			kc:       fake.NewSimpleClientset(),
		},
		{
			name:     "node ready, pods not found",
			nodeName: "toad",
			wantErr:  false,
			want:     false,
			kc: fake.NewSimpleClientset(&corev1.Node{
				TypeMeta: metav1.TypeMeta{
					Kind: "node",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "toad",
				},
				Status: corev1.NodeStatus{
					Conditions: []corev1.NodeCondition{
						{
							Type:   corev1.NodeReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}),
		},
		{
			name:     "node ready, pods ready",
			nodeName: "toad",
			wantErr:  false,
			want:     true,
			kc: fake.NewSimpleClientset(&corev1.Node{
				TypeMeta: metav1.TypeMeta{
					Kind: "node",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "toad",
				},
				Status: corev1.NodeStatus{
					Conditions: []corev1.NodeCondition{
						{
							Type:   corev1.NodeReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}, &corev1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind: "pod",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "etcd-toad",
					Namespace: "kube-system",
				},
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}, &corev1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind: "pod",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "api-toad",
					Namespace: "kube-system",
				},
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}, &corev1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind: "pod",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "controllers-toad",
					Namespace: "kube-system",
				},
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			}),
		},
	}
	for _, tt := range tests {
		got, err := masterIsReady(tt.kc, tt.nodeName)
		if (err != nil) != tt.wantErr {
			t.Errorf("masterIsReady() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if got != tt.want {
			t.Errorf("masterIsReady() = %v, want %v", got, tt.want)
		}
	}
}
