package data_test

import (
	"adolesce/internal/biz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CallbackInfo", func() {

	BeforeEach(func() {})
	AfterEach(func() {})
	It("DeleteByID", func() {
		info := biz.CallbackInfo{
			CallbackId:  "callback_id",
			AppId:       "app_id",
			VerifyToken: "verify_token",
			SecretKey:   "secret_key",
			State:       0,
		}
		Expect(infoRepo.Create(&info)).NotTo(HaveOccurred())
		Expect(infoRepo.DeleteByID(int64(info.ID))).NotTo(HaveOccurred())
	})
})
