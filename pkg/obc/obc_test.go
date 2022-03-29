package obc

import (
	"github.com/kube-object-storage/lib-bucket-provisioner/pkg/provisioner/api"
	nbv1 "github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OBC referenced BucketClass", func() {
	Context("No bucketclassNamespace specified", func() {
		It("should return the BucketClass namespace to be same as NooBaa System namespace", func() {
			systemNS := "test"

			obc := &nbv1.ObjectBucketClaim{}

			Expect(getBucketClassNamespace(obc, nil, systemNS)).To(Equal(systemNS))
		})
	})

	Context("bucketclassNamespace specified", func() {
		It("should return the BucketClass namespace to be same as specified in the obc", func() {
			systemNS := "test"
			bcn := "test1"

			obc := &nbv1.ObjectBucketClaim{
				Spec: nbv1.ObjectBucketClaimSpec{
					AdditionalConfig: map[string]string{
						"bucketclassNamespace": bcn,
					},
				},
			}

			Expect(getBucketClassNamespace(obc, nil, systemNS)).To(Equal(bcn))
		})
	})

	Context("bucketclassNamespace specified in bucketOptions", func() {
		It("should return the BucketClass namespace to be same as specified in the bucketOptions", func() {
			systemNS := "test"
			bcn := "test1"

			obc := &nbv1.ObjectBucketClaim{
				Spec: nbv1.ObjectBucketClaimSpec{},
			}

			bucketOptions := &api.BucketOptions{
				Parameters: map[string]string{
					"bucketclassNamespace": bcn,
				},
			}

			Expect(getBucketClassNamespace(obc, bucketOptions, systemNS)).To(Equal(bcn))
		})
	})
})
