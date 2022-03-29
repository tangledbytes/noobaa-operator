package operator

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	rbacv1 "k8s.io/api/rbac/v1"
)

var _ = Describe("Operator Suite", func() {
	Context("Cluster Role verification", func() {
		It("Must escalate privileges when WATCH_NAMESPACE is \"\"", func() {
			clusterRole := &rbacv1.ClusterRole{
				Rules: []rbacv1.PolicyRule{},
			}

			// Set WATCH_NAMEPACE to ""
			os.Setenv("WATCH_NAMESPACE", "")

			configureClusterRole(clusterRole)

			Expect(clusterRole.Rules).To(HaveLen(3))
			Expect(clusterRole.Rules).To(Equal([]rbacv1.PolicyRule{
				{
					APIGroups: []string{""},
					Resources: []string{"pods", "services", "persistentvolumeclaims", "serviceaccounts"},
					Verbs:     []string{"*"},
				},
				{
					APIGroups: []string{"apps"},
					Resources: []string{"deployments", "statefulsets"},
					Verbs:     []string{"*"},
				},
				{
					APIGroups: []string{"autoscaling"},
					Resources: []string{"horizontalpodautoscalers"},
					Verbs:     []string{"*"},
				},
			}))

			os.Unsetenv("WATCH_NAMESPACE")
		})

		It("Must leave cluster role untouched when privilege escalation is not required", func() {
			clusterRole := &rbacv1.ClusterRole{
				Rules: []rbacv1.PolicyRule{},
			}

			configureClusterRole(clusterRole)

			Expect(clusterRole.Rules).To(HaveLen(0))
		})
	})
})
