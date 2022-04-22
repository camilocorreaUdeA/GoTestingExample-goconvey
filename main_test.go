package main

import (
	"go_convey_testing/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

// Testing with mocks
func TestWithMocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mocks.NewMockCustomClient(ctrl)
	Convey("Testing mock methods", t, func() {
		mockClient.EXPECT().GetBoolResponse(1).Return(true)
		mockClient.EXPECT().GetStringResponse(true).Return("true")
		mockClient.EXPECT().GetntegerResponse("uno").Return(1)
		svc := NewCustomService(mockClient)
		resp := svc.FetchData()
		So(resp, ShouldEqual, "Response: true, true, 1")
	})
}
