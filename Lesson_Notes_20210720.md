### Slice Range Syntax
- example: sample['a','b','c','d'] 
    - sample[start_index_inclusive : up_to_index_not_inclusive]
    - sample[0:2] 
        - starts from 0 index which is 'a' and ends before 2 index which is 'c'
        - returns 'a', 'b' only
    - sample[1:3]
        - starts from 1 index which is 'b' and ends before 3 index which is 'd'
        - returns 'b','c'
    - sample[:3]
        - starts from the beginning index which is 'a' and ends before 2 index which is 'd'
        - returns 'a','b','c'
    - sample[2:]
        - starts from 2 index which is 'c' and ends at the last index which is 'd'
        - returns 'c' , 'd'


### Byte Slice
 - List of bytes that represent UTF-8 encodings of Unicode Code points (See: ASCII TABLE Site)
 - ex) bs := []byte{71, 111} 
       fmt.Printf("%s", bs)
       %s converts byte slice to string 

### Byte Slice Type Conversion
 - Use []byte()
 - ex) greeting := "Hi There!" 
       fmt.Println([]byte(greeting))


 ### REF1 : https://jacking75.github.io/go_slice_array/
 ### REF2 : https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go-8a99012dcc8f