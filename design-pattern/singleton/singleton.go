package singleton

type Singleton interface{
	AddOne () int 
}

type singleton struct (
	count int 
)

var instance *singleton


func GetInstance() Singleton{

	return Singleton(count: 100 )

} 

func (s *singleton) AddOne() int {
	s.count ++
	return s.count
}