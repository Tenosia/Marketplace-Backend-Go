package service

import (
	"github.com/marketplace-go-backend/services/common/genproto/notification"
)

type NotificationService struct {
	errCh chan error
}

type NotificationServiceImpl interface {
	UserVerifyingEmail(receiverEmail, htmlTemplateName, verifyLink string) error
	UserForgotPassword(receiverEmail, htmlTemplateName, resetLink, username string) error
	UserSucessResetPassword(receiverEmail, htmlTemplateName, username string) error
	SendEmailChatNotification(receiverEmail, senderEmail, message string) error
	SellerHasCompletedAnOrder(data *notification.SellerCompletedAnOrderRequest) error
	SellerRequestDeadlineExtension(data *notification.SellerDeadlineExtensionRequest) error
	BuyerDeadlineExtensionResponse(data *notification.BuyerDeadlineExtension) error
	BuyerRefundsAnOrder(data *notification.BuyerRefundsOrderRequest) error
	SellerCanceledAnOrder(data *notification.SellerCancelOrderRequest) error
	NotifySellerGotAnOrder(data *notification.NotifySellerGotAnOrderRequest) error
	NotifySellerGotAReview(data *notification.NotifySellerGotAReviewRequest) error
	NotifyBuyerSellerDeliveredOrder(data *notification.NotifyBuyerOrderDeliveredRequest) error
	NotifyBuyerSellerProcessedOrder(data *notification.NotifyBuyerOrderAcknowledgeRequest) error
}

func NewNotificationService() NotificationServiceImpl {
	return &NotificationService{
		errCh: make(chan error, 1),
	}
}

// TODO: Implement notification service methods
