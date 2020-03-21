package geometry

// cubeの体積計算
func CubeVolume(n int) (int, error) int {
	if n!=0{
		return n * n * n
	}else{
		return 0, errors.New("Zero length edge is not allowed")
	}
	
}

// vMAJOR.MINOR.PATCH
/*
Semantic Versioning = ユニバーサルなバージョニング仕様
x.y.z
x:major
y:minor
z:patch version number
v1.2.3
major1,minor2,patch3

*/
