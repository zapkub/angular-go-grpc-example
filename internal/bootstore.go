package internal

import (
	"context"

	"github.com/zapkub/bootstore-angular-go-grpc/pkg/apis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var bookList = []*apis.BookInfo{
	{
		Id:           StringPtr("bs001"),
		Title:        StringPtr("Lord Of The Ring: Return Of The King"),
		ChapterCount: Int32Ptr(20),
		Price:        Int32Ptr(590),
	},
	{
		Id:           StringPtr("bs002"),
		Title:        StringPtr("Lord Of The Ring: Two Towers"),
		ChapterCount: Int32Ptr(29),
		Price:        Int32Ptr(650),
	},
}
var userInventory = map[string][]*apis.BookInfo{}

type BookStoreEndpointHandler struct {
	apis.UnimplementedBookstoreServer
}

func (b *BookStoreEndpointHandler) GetBookInfoList(ctx context.Context, req *apis.GetBookInfoListRequest) (*apis.GetBookInfoListResponse, error) {
	return &apis.GetBookInfoListResponse{
		BookList: bookList,
	}, nil
}

func (b *BookStoreEndpointHandler) PurchaseBook(ctx context.Context, req *apis.PurchaseBookRequest) (*apis.PurchaseBookResponse, error) {

	var customerName = req.GetCustomerName()
	if userInventory[customerName] == nil {
		userInventory[customerName] = make([]*apis.BookInfo, 0)
	}

	var orderedBook *apis.BookInfo
	for _, book := range bookList {
		if book.GetId() == req.GetBookId() {
			orderedBook = book
		}
	}
	if orderedBook == nil {
		return nil, status.Errorf(codes.NotFound, "book id = %s not exists", req.GetBookId())
	}

	userInventory[customerName] = append(userInventory[customerName], orderedBook)

	return &apis.PurchaseBookResponse{
		OrderId: StringPtr("mock_order_id"),
	}, nil
}
func (b *BookStoreEndpointHandler) MyInventory(ctx context.Context, req *apis.MyInventoryRequest) (*apis.MyInventoryResponse, error) {
	return &apis.MyInventoryResponse{
		BookList: userInventory[req.GetCustomerName()],
	}, nil
}

func Int32Ptr(v int32) *int32 {
	return &v
}
func StringPtr(v string) *string {
	return &v
}
