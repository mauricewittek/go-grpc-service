package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocket(id).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocket(context.Background(), id)
		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}

		if rkt.ID != id {
			t.Errorf("Expected rocket to have id %q, got %q", id, rkt.ID)
		}
	})

	t.Run("test insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(Rocket{
				ID: id,
			}).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(context.Background(), Rocket{ID: id})
		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}

		if rkt.ID != id {
			t.Errorf("Expected rocket to have id %q, got %q", id, rkt.ID)
		}
	})

	t.Run("test delete rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocket(id).
			Return(nil)

		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(context.Background(), id)
		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}
	})
}
