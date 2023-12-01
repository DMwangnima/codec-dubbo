// Code generated by Kitex v0.8.0. DO NOT EDIT.

package testservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error)
	EchoByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error)
	EchoInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error)
	EchoInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error)
	EchoFloat(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error)
	EchoDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error)
	EchoString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error)
	EchoBinary(ctx context.Context, req []byte, callOptions ...callopt.Option) (r []byte, err error)
	Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error)
	EchoBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error)
	EchoByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error)
	EchoInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error)
	EchoInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error)
	EchoInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error)
	EchoFloatList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error)
	EchoDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error)
	EchoStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error)
	EchoBinaryList(ctx context.Context, req [][]byte, callOptions ...callopt.Option) (r [][]byte, err error)
	EchoBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error)
	EchoBool2ByteMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error)
	EchoBool2Int16Map(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error)
	EchoBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error)
	EchoBool2Int64Map(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error)
	EchoBool2FloatMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error)
	EchoBool2DoubleMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error)
	EchoBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error)
	EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte, callOptions ...callopt.Option) (r map[bool][]byte, err error)
	EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool, callOptions ...callopt.Option) (r map[bool][]bool, err error)
	EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8, callOptions ...callopt.Option) (r map[bool][]int8, err error)
	EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16, callOptions ...callopt.Option) (r map[bool][]int16, err error)
	EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32, callOptions ...callopt.Option) (r map[bool][]int32, err error)
	EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64, callOptions ...callopt.Option) (r map[bool][]int64, err error)
	EchoBool2FloatListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error)
	EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error)
	EchoBool2StringListMap(ctx context.Context, req map[bool][]string, callOptions ...callopt.Option) (r map[bool][]string, err error)
	EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte, callOptions ...callopt.Option) (r map[bool][][]byte, err error)
	EchoMultiBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool, callOptions ...callopt.Option) (r *echo.EchoMultiBoolResponse, err error)
	EchoMultiByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8, callOptions ...callopt.Option) (r *echo.EchoMultiByteResponse, err error)
	EchoMultiInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16, callOptions ...callopt.Option) (r *echo.EchoMultiInt16Response, err error)
	EchoMultiInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32, callOptions ...callopt.Option) (r *echo.EchoMultiInt32Response, err error)
	EchoMultiInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64, callOptions ...callopt.Option) (r *echo.EchoMultiInt64Response, err error)
	EchoMultiFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiFloatResponse, err error)
	EchoMultiDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiDoubleResponse, err error)
	EchoMultiString(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string, callOptions ...callopt.Option) (r *echo.EchoMultiStringResponse, err error)
	EchoBaseBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error)
	EchoBaseByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error)
	EchoBaseInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error)
	EchoBaseInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoBaseInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error)
	EchoBaseFloat(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error)
	EchoBaseDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error)
	EchoBaseBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error)
	EchoBaseByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error)
	EchoBaseInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error)
	EchoBaseInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error)
	EchoBaseInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error)
	EchoBaseFloatList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error)
	EchoBaseDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error)
	EchoBool2BoolBaseMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error)
	EchoBool2ByteBaseMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error)
	EchoBool2Int16BaseMap(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error)
	EchoBool2Int32BaseMap(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error)
	EchoBool2Int64BaseMap(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error)
	EchoBool2FloatBaseMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error)
	EchoBool2DoubleBaseMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error)
	EchoMultiBaseBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool, callOptions ...callopt.Option) (r *echo.EchoMultiBoolResponse, err error)
	EchoMultiBaseByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8, callOptions ...callopt.Option) (r *echo.EchoMultiByteResponse, err error)
	EchoMultiBaseInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16, callOptions ...callopt.Option) (r *echo.EchoMultiInt16Response, err error)
	EchoMultiBaseInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32, callOptions ...callopt.Option) (r *echo.EchoMultiInt32Response, err error)
	EchoMultiBaseInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64, callOptions ...callopt.Option) (r *echo.EchoMultiInt64Response, err error)
	EchoMultiBaseFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiFloatResponse, err error)
	EchoMultiBaseDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiDoubleResponse, err error)
	EchoMethodA(ctx context.Context, req bool, callOptions ...callopt.Option) (r string, err error)
	EchoMethodB(ctx context.Context, req int32, callOptions ...callopt.Option) (r string, err error)
	EchoMethodC(ctx context.Context, req int32, callOptions ...callopt.Option) (r string, err error)
	EchoMethodD(ctx context.Context, req1 bool, req2 int32, callOptions ...callopt.Option) (r string, err error)
	EchoOptionalBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error)
	EchoOptionalInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoOptionalString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error)
	EchoOptionalBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error)
	EchoOptionalInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error)
	EchoOptionalStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error)
	EchoOptionalBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error)
	EchoOptionalBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error)
	EchoOptionalBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error)
	EchoOptionalStruct(ctx context.Context, req *echo.EchoOptionalStructRequest, callOptions ...callopt.Option) (r *echo.EchoOptionalStructResponse, err error)
	EchoOptionalMultiBoolRequest(ctx context.Context, req *echo.EchoOptionalMultiBoolRequest, callOptions ...callopt.Option) (r bool, err error)
	EchoOptionalMultiInt32Request(ctx context.Context, req *echo.EchoOptionalMultiInt32Request, callOptions ...callopt.Option) (r int32, err error)
	EchoOptionalMultiStringRequest(ctx context.Context, req *echo.EchoOptionalMultiStringRequest, callOptions ...callopt.Option) (r string, err error)
	EchoOptionalMultiBoolResponse(ctx context.Context, req bool, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiBoolResponse, err error)
	EchoOptionalMultiInt32Response(ctx context.Context, req int32, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiInt32Response, err error)
	EchoOptionalMultiStringResponse(ctx context.Context, req string, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiStringResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTestServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTestServiceClient struct {
	*kClient
}

func (p *kTestServiceClient) EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt(ctx, req)
}

func (p *kTestServiceClient) EchoBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool(ctx, req)
}

func (p *kTestServiceClient) EchoByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoByte(ctx, req)
}

func (p *kTestServiceClient) EchoInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt16(ctx, req)
}

func (p *kTestServiceClient) EchoInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt32(ctx, req)
}

func (p *kTestServiceClient) EchoInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt64(ctx, req)
}

func (p *kTestServiceClient) EchoFloat(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoFloat(ctx, req)
}

func (p *kTestServiceClient) EchoDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoDouble(ctx, req)
}

func (p *kTestServiceClient) EchoString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoString(ctx, req)
}

func (p *kTestServiceClient) EchoBinary(ctx context.Context, req []byte, callOptions ...callopt.Option) (r []byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBinary(ctx, req)
}

func (p *kTestServiceClient) Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}

func (p *kTestServiceClient) EchoBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBoolList(ctx, req)
}

func (p *kTestServiceClient) EchoByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoByteList(ctx, req)
}

func (p *kTestServiceClient) EchoInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt16List(ctx, req)
}

func (p *kTestServiceClient) EchoInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt32List(ctx, req)
}

func (p *kTestServiceClient) EchoInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt64List(ctx, req)
}

func (p *kTestServiceClient) EchoFloatList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoFloatList(ctx, req)
}

func (p *kTestServiceClient) EchoDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoDoubleList(ctx, req)
}

func (p *kTestServiceClient) EchoStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoStringList(ctx, req)
}

func (p *kTestServiceClient) EchoBinaryList(ctx context.Context, req [][]byte, callOptions ...callopt.Option) (r [][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBinaryList(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BoolMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2ByteMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2ByteMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int16Map(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int16Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int32Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int64Map(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int64Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2FloatMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2FloatMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2DoubleMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2DoubleMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2StringMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte, callOptions ...callopt.Option) (r map[bool][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BinaryMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool, callOptions ...callopt.Option) (r map[bool][]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BoolListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8, callOptions ...callopt.Option) (r map[bool][]int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2ByteListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16, callOptions ...callopt.Option) (r map[bool][]int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int16ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32, callOptions ...callopt.Option) (r map[bool][]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int32ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64, callOptions ...callopt.Option) (r map[bool][]int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int64ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2FloatListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2FloatListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2DoubleListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2StringListMap(ctx context.Context, req map[bool][]string, callOptions ...callopt.Option) (r map[bool][]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2StringListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte, callOptions ...callopt.Option) (r map[bool][][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BinaryListMap(ctx, req)
}

func (p *kTestServiceClient) EchoMultiBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool, callOptions ...callopt.Option) (r *echo.EchoMultiBoolResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBool(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8, callOptions ...callopt.Option) (r *echo.EchoMultiByteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiByte(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16, callOptions ...callopt.Option) (r *echo.EchoMultiInt16Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiInt16(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32, callOptions ...callopt.Option) (r *echo.EchoMultiInt32Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiInt32(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64, callOptions ...callopt.Option) (r *echo.EchoMultiInt64Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiInt64(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiFloatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiFloat(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiDoubleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiDouble(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiString(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string, callOptions ...callopt.Option) (r *echo.EchoMultiStringResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiString(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoBaseBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseBool(ctx, req)
}

func (p *kTestServiceClient) EchoBaseByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseByte(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt16(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt32(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt64(ctx, req)
}

func (p *kTestServiceClient) EchoBaseFloat(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseFloat(ctx, req)
}

func (p *kTestServiceClient) EchoBaseDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseDouble(ctx, req)
}

func (p *kTestServiceClient) EchoBaseBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseBoolList(ctx, req)
}

func (p *kTestServiceClient) EchoBaseByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseByteList(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt16List(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt32List(ctx, req)
}

func (p *kTestServiceClient) EchoBaseInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseInt64List(ctx, req)
}

func (p *kTestServiceClient) EchoBaseFloatList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseFloatList(ctx, req)
}

func (p *kTestServiceClient) EchoBaseDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBaseDoubleList(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BoolBaseMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BoolBaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2ByteBaseMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2ByteBaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int16BaseMap(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int16BaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int32BaseMap(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int32BaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int64BaseMap(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int64BaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2FloatBaseMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2FloatBaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2DoubleBaseMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2DoubleBaseMap(ctx, req)
}

func (p *kTestServiceClient) EchoMultiBaseBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool, callOptions ...callopt.Option) (r *echo.EchoMultiBoolResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseBool(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8, callOptions ...callopt.Option) (r *echo.EchoMultiByteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseByte(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16, callOptions ...callopt.Option) (r *echo.EchoMultiInt16Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseInt16(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32, callOptions ...callopt.Option) (r *echo.EchoMultiInt32Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseInt32(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64, callOptions ...callopt.Option) (r *echo.EchoMultiInt64Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseInt64(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiFloatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseFloat(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMultiBaseDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64, callOptions ...callopt.Option) (r *echo.EchoMultiDoubleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMultiBaseDouble(ctx, baseReq, listReq, mapReq)
}

func (p *kTestServiceClient) EchoMethodA(ctx context.Context, req bool, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMethodA(ctx, req)
}

func (p *kTestServiceClient) EchoMethodB(ctx context.Context, req int32, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMethodB(ctx, req)
}

func (p *kTestServiceClient) EchoMethodC(ctx context.Context, req int32, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMethodC(ctx, req)
}

func (p *kTestServiceClient) EchoMethodD(ctx context.Context, req1 bool, req2 int32, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoMethodD(ctx, req1, req2)
}

func (p *kTestServiceClient) EchoOptionalBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalBool(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalInt32(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalString(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalBoolList(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalInt32List(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalStringList(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalBool2BoolMap(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalBool2Int32Map(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalBool2StringMap(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalStruct(ctx context.Context, req *echo.EchoOptionalStructRequest, callOptions ...callopt.Option) (r *echo.EchoOptionalStructResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalStruct(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiBoolRequest(ctx context.Context, req *echo.EchoOptionalMultiBoolRequest, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiBoolRequest(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiInt32Request(ctx context.Context, req *echo.EchoOptionalMultiInt32Request, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiInt32Request(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiStringRequest(ctx context.Context, req *echo.EchoOptionalMultiStringRequest, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiStringRequest(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiBoolResponse(ctx context.Context, req bool, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiBoolResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiBoolResponse(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiInt32Response(ctx context.Context, req int32, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiInt32Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiInt32Response(ctx, req)
}

func (p *kTestServiceClient) EchoOptionalMultiStringResponse(ctx context.Context, req string, callOptions ...callopt.Option) (r *echo.EchoOptionalMultiStringResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOptionalMultiStringResponse(ctx, req)
}
