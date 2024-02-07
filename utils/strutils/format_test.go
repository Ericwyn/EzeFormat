package strutils

import (
	"fmt"
	"testing"
)

func TestGuessType(t *testing.T) {
	//fmt.Println(GuessType("{'aa' : 'bb'}"))
	//fmt.Println(GuessType("<aa>bb</aa>"))
	//fmt.Println(GuessType("1234567890"))
	//fmt.Println(GuessType("1234567890123"))
	//fmt.Println(GuessType("2023-04-03"))
	//fmt.Println(GuessType("2023-04-03 00:00"))
	//fmt.Println(GuessType("2023-04-03 00:00:00"))
	//fmt.Println(GuessType("2023-04-03 00:00:00.000"))

	fmt.Println(GuessType("2023-04-03"))
}

func TestFormatSmart(t *testing.T) {
	fmt.Println(FormatSmart("2023-04-03 18:18"))
}

func TestFormatXml(t *testing.T) {
	xml := formatXml("<?xml version=\"1.0\" encoding=\"UTF-8\"?><createShipmentRequest><integrationHeader><transactionId>CD000100001419</transactionId><applicationId>CIDER23</applicationId><userId>CIDER23</userId><password>gq!2pLk!8d</password></integrationHeader><shipment><shipper><shipperCompanyName>北京荔枝与芒果科技有限公司</shipperCompanyName><shipperAddressLine1>北江大道与亚铝大街交叉口大旺高新技术开发区唯品会物流园2号库4分区</shipperAddressLine1><shipperAddressLine2>null</shipperAddressLine2><shipperCity>肇庆市</shipperCity><shipperCounty>中国</shipperCounty><shipperCountryCode>CN</shipperCountryCode><shipperPostCode>526238</shipperPostCode><shipperContactName>刘伟民</shipperContactName><shipperPhoneNumber>+85230016196</shipperPhoneNumber><shipperEmailAddress>weimin.liu@shopcider.com</shipperEmailAddress><shipperReference>100015550</shipperReference></shipper><destination><destinationAddressLine1>16 Burnside Crescent Blantyre</destinationAddressLine1><destinationCity>Glasgow</destinationCity><destinationCounty>United Kingdom</destinationCounty><destinationCountryCode>GB</destinationCountryCode><destinationPostCode>G72 0LB</destinationPostCode><destinationContactName>Caroline Raggett</destinationContactName><destinationPhoneNumber>07940485964</destinationPhoneNumber><destinationEmailAddress>carolineraggett@yahoo.com</destinationEmailAddress></destination><shipmentInformation><shipmentDate>2023-04-12</shipmentDate><serviceCode>ITLN</serviceCode><serviceOptions><postingLocation>9000257150</postingLocation></serviceOptions><totalPackages>1</totalPackages><totalWeight>0.152</totalWeight><weightId>K</weightId><product>NDX</product><descriptionOfGoods>Long Sleeve Tees</descriptionOfGoods><declaredValue>15.60</declaredValue><declaredCurrencyCode>USD</declaredCurrencyCode><customsInformation><shippingCharges>0</shippingCharges></customsInformation><labelImageFormat>PDF</labelImageFormat><packages><package><packageId>100015550</packageId></package></packages><itemInformation><packageId>100015550</packageId><itemHsCode>6103320000</itemHsCode><itemDescription>Long Sleeve Tees</itemDescription><itemQuantity>1</itemQuantity><itemValue>15.6</itemValue><itemCOO>CN</itemCOO><itemNetWeight>0.152</itemNetWeight></itemInformation></shipmentInformation></shipment></createShipmentRequest>")
	fmt.Println(xml)
	fmt.Println("----------------")
	fmt.Println(compressXml(xml))
}
