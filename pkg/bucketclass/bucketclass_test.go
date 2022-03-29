package bucketclass

import (
	nbv1 "github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Bucketclass compatibility with Namespace store and Backing store", func() {
	Context("BucketClass in the same namespace as store", func() {
		It("should be comptabile", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeTrue())
		})
	})

	Context("BucketClass in different namespace from store: store explicitly permits the namespace", func() {
		It("should be comptabile", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
					Annotations: map[string]string{
						"allowedNamespaces": "test1",
					},
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeTrue())
		})
	})

	Context("BucketClass in different namespace from store: store permits multiple namespaces", func() {
		It("should be comptabile", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
					Annotations: map[string]string{
						"allowedNamespaces": "test1,test2",
					},
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeTrue())
		})

		It("should be comptabile with illformated policy string", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
					Annotations: map[string]string{
						"allowedNamespaces": "test1,test2, test3,",
					},
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeTrue())
		})
	})

	Context("BucketClass in different namespace from store: store does not allow namespace", func() {
		It("should be incomptabile", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
					Annotations: map[string]string{
						"allowedNamespaces": "test2",
					},
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeFalse())
		})
	})

	Context("BucketClass in different namespace from store: store does not specify any policy", func() {
		It("should be incomptabile", func() {
			bc := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
				},
			}

			store := &nbv1.BackingStore{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test",
				},
			}

			Expect(isStoreCompatible(bc, store)).To(BeFalse())
		})
	})
})
