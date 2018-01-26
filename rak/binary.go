package rak

var magic = []byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}

func writeMagic(buf []byte, off *int) {
	for i := 0; i < len(magic); i++ {
		buf[*off] = magic[i]; *off ++
	}
}

func validateMagic(buf []byte, off *int) bool {
	for i := 0; i < len(magic); i++ {
		if buf[*off] != magic[i] {
			return false
		}
		*off ++
	}
	return true
}

func writeInt(buf []byte, off *int, n int32) {
	buf[*off] = byte(n >> 24); 	*off ++
	buf[*off] = byte(n >> 16); 	*off ++
	buf[*off] = byte(n >> 8); 	*off ++
	buf[*off] = byte(n); 		*off ++
}

func readInt(buf []byte, off *int, ans *int32) {
	*ans = 0
	*ans += int32(buf[*off]) << 24; *off ++
	*ans += int32(buf[*off]) << 16; *off ++
	*ans += int32(buf[*off]) << 8; 	*off ++
	*ans += int32(buf[*off]); 		*off ++
}
func writeLong(buf []byte, off *int, n int64) {
	buf[*off] = byte(n >> 56); 	*off ++
	buf[*off] = byte(n >> 48); 	*off ++
	buf[*off] = byte(n >> 40); 	*off ++
	buf[*off] = byte(n >> 32); 	*off ++
	buf[*off] = byte(n >> 24); 	*off ++
	buf[*off] = byte(n >> 16); 	*off ++
	buf[*off] = byte(n >> 8); 	*off ++
	buf[*off] = byte(n); 		*off ++
}

func readLong(buf []byte, off *int, ans *int64) {
	*ans = 0
	*ans += int64(buf[*off]) << 56; *off ++
	*ans += int64(buf[*off]) << 48; *off ++
	*ans += int64(buf[*off]) << 40; *off ++
	*ans += int64(buf[*off]) << 32; *off ++
	*ans += int64(buf[*off]) << 24; *off ++
	*ans += int64(buf[*off]) << 16; *off ++
	*ans += int64(buf[*off]) << 8; 	*off ++
	*ans += int64(buf[*off]); 		*off ++
}
