package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b

}

//변수를 하나 선언한다 타입은 blockchain을 가리키는 포인터
// 위 변수는 대문자(B *blockchain)가 아니기 때문에 해당 패키지 내에서만 접근 가능하다
//이 변수의 instance를 직접 공유하지 않고
// 이 변수의 instance 대신해서 접근시켜주는 function을 생성하는 것이
// Singleton이다
// 이 function을 생성함으로써 다른 패키지에서 blockchain에 어떻게
// 접근할지를 제어할 수 있다

// blockchain을 어떻게 초기화 하고 공유될지 제어 하는 함수
// blockchain의 생성을 제어하기 때문에 중요한 부분이다
