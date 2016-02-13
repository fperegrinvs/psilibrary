package main_test

import (
	//. "psilibrary/server"
	. "psilibrary/server/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Category", func() {
	
	var	(
		cat1 Category
	)

	BeforeEach(func(){
		cat1 = Category{
			ID: 1,
			Name: "Parent",
			ParentId: 0,
		}
	})

	Describe("Creating new categories", func() {
		It("The method should be exposed in the rest API", func(){
			Expect(0).To(Equal(3))
		})
		Context("validation is ok", func(){
			It("Should return the new id", func(){
				Expect(0).To(Equal(3))
			})
		})
		Context("Parent category does not exists", func(){
			It("Should return -1 and an error message", func(){
				Expect(0).To(Equal(3))
			})
		})
	})
	Describe("Listing all categories", func(){
		It("The method should be exposed in the rest API", func(){
			Expect(0).To(Equal(3))
		})
		It("Should list all categories", func(){
			Expect(0).To(Equal(3))
		})
	})
	Describe("Getting a category", func(){
		It("The method should be exposed in the rest API", func(){
			Expect(0).To(Equal(3))
		})
		Context("the category exists", func(){
			It("Should return the category", func(){
				Expect(0).To(Equal(3))
			})
		})
		Context("the category don't exist", func(){
			It("Should return -1", func(){
				Expect(0).To(Equal(3))
			})
		})
	})
	Describe("Updating a category", func(){
		It("The method should be exposed in the rest API", func(){
			Expect(0).To(Equal(3))
		})
		Context("validation is ok", func(){
			It("Should return 0", func(){
				Expect(0).To(Equal(3))
			})
		})
		Context("Parent category does not exists", func(){
			It("Should return -1 and an error message", func(){
				Expect(0).To(Equal(3))
			})
		})
	})
	Describe("Deleting a category", func(){
		It("The method should be exposed in the rest API", func(){
			Expect(0).To(Equal(3))
		})
		Context("If its used by someone", func(){
			It("Should list all objects depending on it and a keep the category", func(){
				Expect(0).To(Equal(3))
			})
		})
		Context("if its not used by anyone", func(){
			It("Should delete the category and return the id of the deleted object", func(){
				Expect(0).To(Equal(3))
			})
		})
	})
})
