package inter

//ByteConter conter the bytes of the para
type ByteConter int 

func (p *ByteConter) Write(n []byte) (int,error) {
	 *p += ByteConter(len(n))
	 return len(n),nil
}
