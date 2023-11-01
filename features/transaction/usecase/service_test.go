package usecase

import (
	"errors"
	"fmt"
	"perpustakaan/features/member"
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"
	"perpustakaan/features/transaction/mocks"
	helperMocks "perpustakaan/helpers/mocks"
	"testing"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var transactions = []transaction.Transaction{
		{
			ID: 1,
			OrderID: "LOAN-E3E76882-894D-4C0F-A373-CBC5112E32D2", 
			Note: "Pembayaran Denda Kehilangan Buku",
			Status: "Pending",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/d8212007-74a1-4f49-9d21-b18f36324022", 
			MemberID: 1, 
		},
	}

	var page = 1
	var size = 10

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size).Return(transactions, nil).Once()

		result, message := service.FindAll(page, size)
		assert.Equal(t, transactions[0].Note, result[0].Note)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindAll(page, size)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var transaction = transaction.Transaction{
		ID: 1,
		OrderID: "LOAN-E3E76882-894D-4C0F-A373-CBC5112E32D2", 
		Note: "Pembayaran Denda Kehilangan Buku",
		Status: "Pending",
		PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/d8212007-74a1-4f49-9d21-b18f36324022", 
		MemberID: 1, 
	}

	var fineItems = []dtos.FineItem{
		{
			ID: 1,
			Name: "Dark Gathering",
			Status: "Lost",
			Amount: 1,
		},
	}

	var transactionID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", transactionID).Return(&transaction, nil).Once()
		repository.On("SelectAllFineItemOnTransactionID", transactionID).Return(fineItems, nil).Once()

		result, message := service.FindByID(transactionID)
		assert.Equal(t, transaction.Note, result.Note)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Transaction Not Found", func(t *testing.T) {
		repository.On("SelectByID", 0).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(0)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Fine Items Not Found", func(t *testing.T) {
		repository.On("SelectByID", 3).Return(&transaction, nil).Once()
		repository.On("SelectAllFineItemOnTransactionID", 3).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(3)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var transactionID = 1
	var memberID = 1
	var orderID = fmt.Sprintf("LOAN-%s%d%d", "080000000000", 1, 1)

	var validTransaction = transaction.Transaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		OrderID: orderID,
		Status: "Pending",
		PaymentURL: "example.sandbox-payment-midtrands.com",
		MemberID: 1, 
	}

	var invalidTransaction = transaction.Transaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		OrderID: orderID,
		Status: "Pending",
		PaymentURL: "example.sandbox-payment-midtrands.com",
		MemberID: 3,
	}

	var input = dtos.InputTransaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		MemberID: 1,
	}

	var inputV2 = dtos.InputTransaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		MemberID: 4,
		LoanIDS: []int{1},
	}

	var member = member.Member{
		ID: 1,
		FullName: "Nibras",
		CredentialNumber: "2107411026",
		Email: "nibras@example.com",
		Password: "randomgeneratedhash",
		PhoneNumber: "080000000000",
		Address: "jalan hj raisin",
	}
	
	var fineItems = []dtos.FineItem{
		{
			ID: 1,
			Name: "Dark Gathering",
			Status: "Lost",
			Amount: 50000,
		},
		{
			ID: 2,
			Name: "DR Stone",
			Status: "Damaged",
			Amount: 30000,
		},
	}


	var midtransItems = []midtrans.ItemDetails{
		{
			ID: "BOOK-1",
			Name: "Dark Gathering",
			Price: 50000,
			Qty: 1,
		},
		{
			ID: "BOOK-2",
			Name: "DR Stone",
			Price: 30000,
			Qty: 1,
		},
	}

	var customerDetails = midtrans.CustomerDetails{
		FName: member.FullName,
		Email: member.Email,
		Phone: member.PhoneNumber,
	}

	var snapResponse = snap.Response{
		Token: "example-token",
		RedirectURL: "example.sandbox-payment-midtrands.com",
	}

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectMemberByID", memberID).Return(&member, nil).Once()
		repository.On("SelectAllFineItemOnMemberID", memberID).Return(fineItems, nil).Once()
		helper.On("CreatePaymentLink", orderID, int64(80000), midtransItems, customerDetails).Return(&snapResponse, nil).Once()
		repository.On("Insert", validTransaction).Return(1, nil).Once()
		repository.On("UpdateBatchTransactionDetail", fineItems, transactionID).Return(true, nil).Once()

		result, message := service.Create(input)
		assert.Equal(t, result.Note, validTransaction.Note)
		assert.NotNil(t, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed: Error When Insert Transaction", func(t *testing.T) {
		input.MemberID = 2
		invalidTransaction.MemberID = 2
		repository.On("SelectMemberByID", input.MemberID).Return(&member, nil).Once()
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(fineItems, nil).Once()
		helper.On("CreatePaymentLink", orderID, int64(80000), midtransItems, customerDetails).Return(&snapResponse, nil).Once()
		repository.On("Insert", invalidTransaction).Return(0, errors.New("error when insert")).Once()

		result, message := service.Create(input)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed: Error When Create Payment Link", func(t *testing.T) {
		repository.On("SelectMemberByID", input.MemberID).Return(&member, nil).Once()
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(fineItems, nil).Once()
		helper.On("CreatePaymentLink", orderID, int64(80000), midtransItems, customerDetails).Return(nil, errors.New("error when create")).Once()

		result, message := service.Create(input)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed: Error Not Found Fine Item By Member ID", func(t *testing.T) {
		input.MemberID = 3
		repository.On("SelectMemberByID", input.MemberID).Return(&member, nil).Once()
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.Create(input)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		assert.Equal(t, message, "Member does not have any Fines yet!")
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Update Transaction ID On Loan History", func(t *testing.T) {
		validTransaction.ID = 2
		validTransaction.MemberID = 4
		repository.On("SelectMemberByID", inputV2.MemberID).Return(&member, nil).Once()
		repository.On("SelectFineItemByIDAndMemberID", 1, inputV2.MemberID).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.Create(inputV2)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var validTransaction = transaction.Transaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		OrderID: "LOAN-1239JDOIW-123123K-WAD1IO2J",
		Status: "Pending",
		PaymentURL: "example.sandbox-payment-midtrands.com",
		MemberID: 1, 
	}

	var invalidTransaction = transaction.Transaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		OrderID: "LOAN-1239JDOIW-123123K-WAD1IO2J",
		Status: "Pending",
		PaymentURL: "example.sandbox-payment-midtrands.com",
		MemberID: 3,
	}

	var input = dtos.InputTransaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		MemberID: 1,
	}
	
	var inputV2 = dtos.InputTransaction{
		Note: "Pembayaran Denda Kehilangan Buku",
		MemberID: 3,
		LoanIDS: []int{0},
	}
	
	var fineItems = []dtos.FineItem{
		{
			ID: 1,
			Name: "Dark Gathering",
			Status: "Lost",
			Amount: 1,
		},
		{
			ID: 2,
			Name: "DR Stone",
			Status: "Damaged",
			Amount: 1,
		},
	}

	var transactionID = 1
	var memberID = 1
	var orderID = "LOAN-1239JDOIW-123123K-WAD1IO2J"
	var status = "Pending"
	var paymentURL = "example.sandbox-payment-midtrands.com"

	t.Run("Success", func(t *testing.T) {
		validTransaction.ID = transactionID
		repository.On("SelectAllFineItemOnMemberID", memberID).Return(fineItems, nil).Once()
		repository.On("Update", validTransaction).Return(1, nil).Once()
		repository.On("UpdateBatchTransactionDetail", fineItems, transactionID).Return(true, nil).Once()

		result, message := service.Modify(input, transactionID, orderID, status, paymentURL)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed: Error When Update Transaction", func(t *testing.T) {
		input.MemberID = 2
		invalidTransaction.MemberID = 2
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(fineItems, nil).Once()
		repository.On("Update", invalidTransaction).Return(0, errors.New("error when update transaction")).Once()

		result, message := service.Modify(input, 0, orderID, status, paymentURL)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed: Error Unknown Member ID", func(t *testing.T) {
		input.MemberID = 3
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(nil, errors.New("record not found")).Once()

		result, message := service.Modify(input, 0, orderID, status, paymentURL)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Update Transaction ID On Loan History", func(t *testing.T) {
		validTransaction.ID = 2
		input.MemberID =  4
		validTransaction.MemberID = 4
		repository.On("SelectAllFineItemOnMemberID", input.MemberID).Return(fineItems, nil).Once()
		repository.On("Update", validTransaction).Return(1, nil).Once()
		repository.On("UpdateBatchTransactionDetail", fineItems, validTransaction.ID).Return(false, errors.New("error when update")).Once()

		result, message := service.Modify(input, validTransaction.ID, orderID, status, paymentURL)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Unset Transaction ID On Loan History", func(t *testing.T) {
		repository.On("UnsetTransactionIDs", transactionID).Return(false, errors.New("record not found"))

		result, message := service.Modify(inputV2, transactionID, orderID, status, paymentURL)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var transactionID = 1
	
	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", transactionID).Return(1, nil).Once()

		result, message := service.Remove(transactionID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteByID", 0).Return(0, errors.New("record not found")).Once()

		result, message := service.Remove(0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestVerifyPayment(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var orderID = "LOAN-1239JDOIW-123123K-WAD1IO2J"

	var payload = map[string]any{
		"order_id": orderID,
	}

	var invalidPayload = map[string]any{}

	var transactionID = 1

	var transaction = transaction.Transaction{
		ID: transactionID,
		Note: "Pembayaran Denda Kehilangan Buku",
		OrderID: orderID,
		Status: "Pending",
		PaymentURL: "example.sandbox-payment-midtrands.com",
		MemberID: 1, 
	}

	var status = "Success"


	t.Run("Success", func(t *testing.T) {
		helper.On("CheckTransaction", orderID).Return("Success", nil).Once()
		repository.On("SelectTransactionByOrderID", orderID).Return(&transaction, nil).Once()
		repository.On("UpdateStatus", transactionID, status).Return(true, nil).Once()

		result, message := service.VerifyPayment(payload)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Invalid Notification", func(t *testing.T) {
		result, message := service.VerifyPayment(invalidPayload)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error Checking Transaction By Third Party", func(t *testing.T) {
		helper.On("CheckTransaction", orderID).Return("", errors.New("error status")).Once()

		result, message := service.VerifyPayment(payload)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Transaction Not Found", func(t *testing.T) {
		orderID = "LOAN-WADWAF-12ASDS23123K-WAD1IO2J"
		payload["order_id"] = orderID
		helper.On("CheckTransaction", orderID).Return("Success", nil).Once()
		repository.On("SelectTransactionByOrderID", orderID).Return(nil, errors.New("record not found")).Once()

		result, message := service.VerifyPayment(payload)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Update Status Transaction", func(t *testing.T) {
		orderID = "LOAN-WA22WWW2DWAF-12ASDS23123K-WAD1IO2J"
		payload["order_id"] = orderID
		helper.On("CheckTransaction", orderID).Return("Success", nil).Once()
		repository.On("SelectTransactionByOrderID", orderID).Return(&transaction, nil).Once()
		repository.On("UpdateStatus", transactionID, status).Return(false, errors.New("error when update")).Once()

		result, message := service.VerifyPayment(payload)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}