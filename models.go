package retaicrm

//{GET ORDERS

type OrdersRequest struct {
	Filter OrdersFilter `url:"filter,omitempty"`
	Limit  int          `url:"limit,omitempty"`
	Page   int          `url:"page,omitempty"`
}

type OrdersFilter struct {
	Ids                            []int                  `url:"ids,omitempty,brackets"`
	ExternalIds                    []string               `url:"externalIds,omitempty,brackets"`
	Numbers                        []string               `url:"numbers,omitempty,brackets"`
	Customer                       string                 `url:"customer,omitempty"`
	CustomerID                     string                 `url:"customerId,omitempty"`
	CustomerExternalID             string                 `url:"customerExternalId,omitempty"`
	Countries                      []string               `url:"countries,omitempty,brackets"`
	City                           string                 `url:"city,omitempty"`
	Region                         string                 `url:"region,omitempty"`
	Index                          string                 `url:"index,omitempty"`
	Metro                          string                 `url:"metro,omitempty"`
	Email                          string                 `url:"email,omitempty"`
	DeliveryTimeFrom               string                 `url:"deliveryTimeFrom,omitempty"`
	DeliveryTimeTo                 string                 `url:"deliveryTimeTo,omitempty"`
	MinPrepaySumm                  string                 `url:"minPrepaySumm,omitempty"`
	MaxPrepaySumm                  string                 `url:"maxPrepaySumm,omitempty"`
	MinPrice                       string                 `url:"minPrice,omitempty"`
	MaxPrice                       string                 `url:"maxPrice,omitempty"`
	Product                        string                 `url:"product,omitempty"`
	Vip                            int                    `url:"vip,omitempty"`
	Bad                            int                    `url:"bad,omitempty"`
	Attachments                    int                    `url:"attachments,omitempty"`
	Expired                        int                    `url:"expired,omitempty"`
	Call                           int                    `url:"call,omitempty"`
	Online                         int                    `url:"online,omitempty"`
	Shipped                        int                    `url:"shipped,omitempty"`
	UploadedToExtStoreSys          int                    `url:"uploadedToExtStoreSys,omitempty"`
	ReceiptFiscalDocumentAttribute int                    `url:"receiptFiscalDocumentAttribute,omitempty"`
	ReceiptStatus                  int                    `url:"receiptStatus,omitempty"`
	ReceiptOperation               int                    `url:"receiptOperation,omitempty"`
	MinDeliveryCost                string                 `url:"minDeliveryCost,omitempty"`
	MaxDeliveryCost                string                 `url:"maxDeliveryCost,omitempty"`
	MinDeliveryNetCost             string                 `url:"minDeliveryNetCost,omitempty"`
	MaxDeliveryNetCost             string                 `url:"maxDeliveryNetCost,omitempty"`
	ManagerComment                 string                 `url:"managerComment,omitempty"`
	CustomerComment                string                 `url:"customerComment,omitempty"`
	MinMarginSumm                  string                 `url:"minMarginSumm,omitempty"`
	MaxMarginSumm                  string                 `url:"maxMarginSumm,omitempty"`
	MinPurchaseSumm                string                 `url:"minPurchaseSumm,omitempty"`
	MaxPurchaseSumm                string                 `url:"maxPurchaseSumm,omitempty"`
	MinCostSumm                    string                 `url:"minCostSumm,omitempty"`
	MaxCostSumm                    string                 `url:"maxCostSumm,omitempty"`
	TrackNumber                    string                 `url:"trackNumber,omitempty"`
	ContragentName                 string                 `url:"contragentName,omitempty"`
	ContragentInn                  string                 `url:"contragentInn,omitempty"`
	ContragentKpp                  string                 `url:"contragentKpp,omitempty"`
	ContragentBik                  string                 `url:"contragentBik,omitempty"`
	ContragentCorrAccount          string                 `url:"contragentCorrAccount,omitempty"`
	ContragentBankAccount          string                 `url:"contragentBankAccount,omitempty"`
	ContragentTypes                []string               `url:"contragentTypes,omitempty,brackets"`
	OrderTypes                     []string               `url:"orderTypes,omitempty,brackets"`
	PaymentStatuses                []string               `url:"paymentStatuses,omitempty,brackets"`
	PaymentTypes                   []string               `url:"paymentTypes,omitempty,brackets"`
	DeliveryTypes                  []string               `url:"deliveryTypes,omitempty,brackets"`
	OrderMethods                   []string               `url:"orderMethods,omitempty,brackets"`
	ShipmentStores                 []string               `url:"shipmentStores,omitempty,brackets"`
	Couriers                       []string               `url:"couriers,omitempty,brackets"`
	Managers                       []string               `url:"managers,omitempty,brackets"`
	ManagerGroups                  []string               `url:"managerGroups,omitempty,brackets"`
	Sites                          []string               `url:"sites,omitempty,brackets"`
	CreatedAtFrom                  string                 `url:"createdAtFrom,omitempty"`
	CreatedAtTo                    string                 `url:"createdAtTo,omitempty"`
	FullPaidAtFrom                 string                 `url:"fullPaidAtFrom,omitempty"`
	FullPaidAtTo                   string                 `url:"fullPaidAtTo,omitempty"`
	DeliveryDateFrom               string                 `url:"deliveryDateFrom,omitempty"`
	DeliveryDateTo                 string                 `url:"deliveryDateTo,omitempty"`
	StatusUpdatedAtFrom            string                 `url:"statusUpdatedAtFrom,omitempty"`
	StatusUpdatedAtTo              string                 `url:"statusUpdatedAtTo,omitempty"`
	DpdParcelDateFrom              string                 `url:"dpdParcelDateFrom,omitempty"`
	DpdParcelDateTo                string                 `url:"dpdParcelDateTo,omitempty"`
	FirstWebVisitFrom              string                 `url:"firstWebVisitFrom,omitempty"`
	FirstWebVisitTo                string                 `url:"firstWebVisitTo,omitempty"`
	LastWebVisitFrom               string                 `url:"lastWebVisitFrom,omitempty"`
	LastWebVisitTo                 string                 `url:"lastWebVisitTo,omitempty"`
	FirstOrderFrom                 string                 `url:"firstOrderFrom,omitempty"`
	FirstOrderTo                   string                 `url:"firstOrderTo,omitempty"`
	LastOrderFrom                  string                 `url:"lastOrderFrom,omitempty"`
	LastOrderTo                    string                 `url:"lastOrderTo,omitempty"`
	ShipmentDateFrom               string                 `url:"shipmentDateFrom,omitempty"`
	ShipmentDateTo                 string                 `url:"shipmentDateTo,omitempty"`
	ExtendedStatus                 []string               `url:"extendedStatus,omitempty,brackets"`
	SourceName                     string                 `url:"sourceName,omitempty"`
	MediumName                     string                 `url:"mediumName,omitempty"`
	CampaignName                   string                 `url:"campaignName,omitempty"`
	KeywordName                    string                 `url:"keywordName,omitempty"`
	AdContentName                  string                 `url:"adContentName,omitempty"`
	CustomFields                   map[string]interface{} `url:"customFields,omitempty,brackets"`
}

type OrdersResponse struct {
	Success    bool        `json:"success"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Orders     []Order     `json:"orders,omitempty"`
}

type Order struct {
	ID                            int                    `json:"id,omitempty"`
	ExternalID                    string                 `json:"externalId,omitempty"`
	Number                        string                 `json:"number,omitempty"`
	FirstName                     string                 `json:"firstName,omitempty"`
	LastName                      string                 `json:"lastName,omitempty"`
	Patronymic                    string                 `json:"patronymic,omitempty"`
	Email                         string                 `json:"email,omitempty"`
	Phone                         string                 `json:"phone,omitempty"`
	AdditionalPhone               string                 `json:"additionalPhone,omitempty"`
	CreatedAt                     string                 `json:"createdAt,omitempty"`
	StatusUpdatedAt               string                 `json:"statusUpdatedAt,omitempty"`
	ManagerID                     int                    `json:"managerId,omitempty"`
	Mark                          int                    `json:"mark,omitempty"`
	Call                          bool                   `json:"call,omitempty"`
	Expired                       bool                   `json:"expired,omitempty"`
	FromAPI                       bool                   `json:"fromApi,omitempty"`
	MarkDatetime                  string                 `json:"markDatetime,omitempty"`
	CustomerComment               string                 `json:"customerComment,omitempty"`
	ManagerComment                string                 `json:"managerComment,omitempty"`
	Status                        string                 `json:"status,omitempty"`
	StatusComment                 string                 `json:"statusComment,omitempty"`
	FullPaidAt                    string                 `json:"fullPaidAt,omitempty"`
	Site                          string                 `json:"site,omitempty"`
	OrderType                     string                 `json:"orderType,omitempty"`
	OrderMethod                   string                 `json:"orderMethod,omitempty"`
	CountryIso                    string                 `json:"countryIso,omitempty"`
	Summ                          float32                `json:"summ,omitempty"`
	TotalSumm                     float32                `json:"totalSumm,omitempty"`
	PrepaySum                     float32                `json:"prepaySum,omitempty"`
	PurchaseSumm                  float32                `json:"purchaseSumm,omitempty"`
	DiscountManualAmount          float32                `json:"discountManualAmount,omitempty"`
	DiscountManualPercent         float32                `json:"discountManualPercent,omitempty"`
	Weight                        float32                `json:"weight,omitempty"`
	Length                        int                    `json:"length,omitempty"`
	Width                         int                    `json:"width,omitempty"`
	Height                        int                    `json:"height,omitempty"`
	ShipmentStore                 string                 `json:"shipmentStore,omitempty"`
	ShipmentDate                  string                 `json:"shipmentDate,omitempty"`
	ClientID                      string                 `json:"clientId,omitempty"`
	Shipped                       bool                   `json:"shipped,omitempty"`
	UploadedToExternalStoreSystem bool                   `json:"uploadedToExternalStoreSystem,omitempty"`
	Source                        *Source                `json:"source,omitempty"`
	Contragent                    *Contragent            `json:"contragent,omitempty"`
	Customer                      *Customer              `json:"customer,omitempty"`
	Delivery                      *OrderDelivery         `json:"delivery,omitempty"`
	Marketplace                   *OrderMarketplace      `json:"marketplace,omitempty"`
	Items                         []OrderItem            `json:"items,omitempty"`
	CustomFields                  map[string]interface{} `json:"customFields,omitempty"`
	Payments                      OrderPayments          `json:"payments,omitempty"`
}

//GET ORDERS}

//{CREATE ORDER

type OrderCreateResponse struct {
	CreateResponse
	Order    Order  `json:"order,omitempty"`
	ErrorMsg string `json:"errorMsg"`
}

type CreateResponse struct {
	Success bool `json:"success"`
	ID      int  `json:"id,omitempty"`
}

//CREATE ORDER}

//ORDERS}

//{GENERAL STRUCT

type Source struct {
	Source   string `json:"source,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Campaign string `json:"campaign,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Contragent struct {
	ContragentType    string `json:"contragentType,omitempty"`
	LegalName         string `json:"legalName,omitempty"`
	LegalAddress      string `json:"legalAddress,omitempty"`
	INN               string `json:"INN,omitempty"`
	OKPO              string `json:"OKPO,omitempty"`
	KPP               string `json:"KPP,omitempty"`
	OGRN              string `json:"OGRN,omitempty"`
	OGRNIP            string `json:"OGRNIP,omitempty"`
	CertificateNumber string `json:"certificateNumber,omitempty"`
	CertificateDate   string `json:"certificateDate,omitempty"`
	BIK               string `json:"BIK,omitempty"`
	Bank              string `json:"bank,omitempty"`
	BankAddress       string `json:"bankAddress,omitempty"`
	CorrAccount       string `json:"corrAccount,omitempty"`
	BankAccount       string `json:"bankAccount,omitempty"`
}
type Customer struct {
	ID                           int                    `json:"id,omitempty"`
	ExternalID                   string                 `json:"externalId,omitempty"`
	FirstName                    string                 `json:"firstName,omitempty"`
	LastName                     string                 `json:"lastName,omitempty"`
	Patronymic                   string                 `json:"patronymic,omitempty"`
	Sex                          string                 `json:"sex,omitempty"`
	Email                        string                 `json:"email,omitempty"`
	Phones                       []Phone                `json:"phones,omitempty"`
	Address                      *Address               `json:"address,omitempty"`
	CreatedAt                    string                 `json:"createdAt,omitempty"`
	Birthday                     string                 `json:"birthday,omitempty"`
	ManagerID                    int                    `json:"managerId,omitempty"`
	Vip                          bool                   `json:"vip,omitempty"`
	Bad                          bool                   `json:"bad,omitempty"`
	Site                         string                 `json:"site,omitempty"`
	Source                       *Source                `json:"source,omitempty"`
	Contragent                   *Contragent            `json:"contragent,omitempty"`
	PersonalDiscount             float32                `json:"personalDiscount,omitempty"`
	CumulativeDiscount           float32                `json:"cumulativeDiscount,omitempty"`
	DiscountCardNumber           string                 `json:"discountCardNumber,omitempty"`
	EmailMarketingUnsubscribedAt string                 `json:"emailMarketingUnsubscribedAt,omitempty"`
	AvgMarginSumm                float32                `json:"avgMarginSumm,omitempty"`
	MarginSumm                   float32                `json:"marginSumm,omitempty"`
	TotalSumm                    float32                `json:"totalSumm,omitempty"`
	AverageSumm                  float32                `json:"averageSumm,omitempty"`
	OrdersCount                  int                    `json:"ordersCount,omitempty"`
	CostSumm                     float32                `json:"costSumm,omitempty"`
	MaturationTime               int                    `json:"maturationTime,omitempty"`
	FirstClientID                string                 `json:"firstClientId,omitempty"`
	LastClientID                 string                 `json:"lastClientId,omitempty"`
	BrowserID                    string                 `json:"browserId,omitempty"`
	MgCustomerID                 string                 `json:"mgCustomerId,omitempty"`
	PhotoURL                     string                 `json:"photoUrl,omitempty"`
	CustomFields                 map[string]interface{} `json:"customFields,omitempty"`
	Tags                         []Tag                  `json:"tags,omitempty"`
}

type Phone struct {
	Number string `json:"number,omitempty"`
}

type OrderDelivery struct {
	Code            string                `json:"code,omitempty"`
	IntegrationCode string                `json:"integrationCode,omitempty"`
	Cost            float32               `json:"cost,omitempty"`
	NetCost         float32               `json:"netCost,omitempty"`
	VatRate         string                `json:"vatRate,omitempty"`
	Date            string                `json:"date,omitempty"`
	Time            *OrderDeliveryTime    `json:"time,omitempty"`
	Address         *Address              `json:"address,omitempty"`
	Service         *OrderDeliveryService `json:"service,omitempty"`
	Data            *OrderDeliveryData    `json:"data,omitempty"`
}

type OrderDeliveryTime struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Custom string `json:"custom,omitempty"`
}

type Address struct {
	Index      string `json:"index,omitempty"`
	CountryIso string `json:"countryIso,omitempty"`
	Region     string `json:"region,omitempty"`
	RegionID   int    `json:"regionId,omitempty"`
	City       string `json:"city,omitempty"`
	CityID     int    `json:"cityId,omitempty"`
	CityType   string `json:"cityType,omitempty"`
	Street     string `json:"street,omitempty"`
	StreetID   int    `json:"streetId,omitempty"`
	StreetType string `json:"streetType,omitempty"`
	Building   string `json:"building,omitempty"`
	Flat       string `json:"flat,omitempty"`
	Floor      int    `json:"floor,omitempty"`
	Block      int    `json:"block,omitempty"`
	House      string `json:"house,omitempty"`
	Housing    string `json:"housing,omitempty"`
	Metro      string `json:"metro,omitempty"`
	Notes      string `json:"notes,omitempty"`
	Text       string `json:"text,omitempty"`
}

type OrderDeliveryService struct {
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Active bool   `json:"active,omitempty"`
}

type OrderDeliveryData struct {
	OrderDeliveryDataBasic
	AdditionalFields map[string]interface{}
}

type OrderDeliveryDataBasic struct {
	TrackNumber        string `json:"trackNumber,omitempty"`
	Status             string `json:"status,omitempty"`
	PickuppointAddress string `json:"pickuppointAddress,omitempty"`
	PayerType          string `json:"payerType,omitempty"`
}

type Tag struct {
	Name     string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
	Attached bool   `json:"attached,omitempty"`
}

type OrderMarketplace struct {
	Code    string `json:"code,omitempty"`
	OrderID string `json:"orderId,omitempty"`
}

type OrderItem struct {
	ID                    int        `json:"id,omitempty"`
	InitialPrice          float32    `json:"initialPrice,omitempty"`
	PurchasePrice         float32    `json:"purchasePrice,omitempty"`
	DiscountTotal         float32    `json:"discountTotal,omitempty"`
	DiscountManualAmount  float32    `json:"discountManualAmount,omitempty"`
	DiscountManualPercent float32    `json:"discountManualPercent,omitempty"`
	ProductName           string     `json:"productName,omitempty"`
	VatRate               string     `json:"vatRate,omitempty"`
	CreatedAt             string     `json:"createdAt,omitempty"`
	Quantity              float32    `json:"quantity,omitempty"`
	Status                string     `json:"status,omitempty"`
	Comment               string     `json:"comment,omitempty"`
	IsCanceled            bool       `json:"isCanceled,omitempty"`
	Offer                 Offer      `json:"offer,omitempty"`
	Properties            Properties `json:"properties,omitempty"`
	PriceType             *PriceType `json:"priceType,omitempty"`
}

type Offer struct {
	ID            int                    `json:"id,omitempty"`
	ExternalID    string                 `json:"externalId,omitempty"`
	Name          string                 `json:"name,omitempty"`
	XMLID         string                 `json:"xmlId,omitempty"`
	Article       string                 `json:"article,omitempty"`
	VatRate       string                 `json:"vatRate,omitempty"`
	Price         float32                `json:"price,omitempty"`
	PurchasePrice float32                `json:"purchasePrice,omitempty"`
	Quantity      float32                `json:"quantity,omitempty"`
	Height        float32                `json:"height,omitempty"`
	Width         float32                `json:"width,omitempty"`
	Length        float32                `json:"length,omitempty"`
	Weight        float32                `json:"weight,omitempty"`
	Stores        []Inventory            `json:"stores,omitempty"`
	Properties    map[string]interface{} `json:"properties,omitempty"`
	Prices        []OfferPrice           `json:"prices,omitempty"`
	Images        []string               `json:"images,omitempty"`
	Unit          *Unit                  `json:"unit,omitempty"`
}

type Inventory struct {
	PurchasePrice float32 `json:"purchasePrice,omitempty"`
	Quantity      float32 `json:"quantity,omitempty"`
	Store         string  `json:"store,omitempty"`
}

type OfferPrice struct {
	Price     float32 `json:"price,omitempty"`
	Ordering  int     `json:"ordering,omitempty"`
	PriceType string  `json:"priceType,omitempty"`
}
type Unit struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Sym     string `json:"sym"`
	Default bool   `json:"default,omitempty"`
	Active  bool   `json:"active,omitempty"`
}

type Properties map[string]Property

type Property struct {
	Code  string   `json:"code,omitempty"`
	Name  string   `json:"name,omitempty"`
	Value string   `json:"value,omitempty"`
	Sites []string `json:"Sites,omitempty"`
}

type PriceType struct {
	ID               int               `json:"id,omitempty"`
	Code             string            `json:"code,omitempty"`
	Name             string            `json:"name,omitempty"`
	Active           bool              `json:"active,omitempty"`
	Default          bool              `json:"default,omitempty"`
	Description      string            `json:"description,omitempty"`
	FilterExpression string            `json:"filterExpression,omitempty"`
	Ordering         int               `json:"ordering,omitempty"`
	Groups           []string          `json:"groups,omitempty"`
	Geo              []GeoHierarchyRow `json:"geo,omitempty"`
}

type GeoHierarchyRow struct {
	Country  string `json:"country,omitempty"`
	Region   string `json:"region,omitempty"`
	RegionID int    `json:"regionId,omitempty"`
	City     string `json:"city,omitempty"`
	CityID   int    `json:"cityId,omitempty"`
}

type OrderPayments map[string]OrderPayment
type OrderPayment struct {
	ID         int     `json:"id,omitempty"`
	ExternalID string  `json:"externalId,omitempty"`
	Type       string  `json:"type,omitempty"`
	Status     string  `json:"status,omitempty"`
	PaidAt     string  `json:"paidAt,omitempty"`
	Amount     float32 `json:"amount,omitempty"`
	Comment    string  `json:"comment,omitempty"`
}

type Pagination struct {
	Limit          int `json:"limit,omitempty"`
	TotalCount     int `json:"totalCount,omitempty"`
	CurrentPage    int `json:"currentPage,omitempty"`
	TotalPageCount int `json:"totalPageCount,omitempty"`
}

//GENERAL STRUCT}
