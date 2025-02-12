// working with files
// Files in go can be read and written using the os and ioutil standard libraries
package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	// Create the file (it will be empty)
	f, err := os.Create("output.txt")
	if err != nil {
		panic("unable to create file")
	}
	defer f.Close()

	// Write something into the file for reading (optional, but this ensures you have content to read)
	_, err = f.Write([]byte("Hello, Go!"))
	if err != nil {
		panic("unable to write to file")
	}

	// Go back to the beginning of the file to read
	f.Seek(0, io.SeekStart)

	// Create a buffer to read data into
	buf := make([]byte, 64)
	cnt, err := f.Read(buf)

	// Check if we've reached the end of the file or encountered an error
	if err != nil && err != io.EOF {
		panic("unable to read file")
	}

	// Print how many bytes we read and the content
	fmt.Printf("Read %d bytes\n", cnt)
	fmt.Println(string(buf[:cnt])) // Converts the byte slice to a string and prints it
}


// 1. Importing Necessary Packages:
// go
// Copy
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )
// The fmt package is used for formatted I/O (printing to the console).
// The os package is used to work with the operating system, specifically for file handling.
// The strings package is imported, but it's not used in this specific code snippet. It might be left over from some previous version or an intended future use.
// 2. Creating the File:
// go
// Copy
// f, err := os.Create("output.txt")
// The os.Create("output.txt") function is called to create a new file named output.txt.
// It returns two values:
// f (a file object): A pointer to the newly created file, which will be used for further file operations.
// err (an error): An error object, which will be non-nil if there is an issue creating the file (e.g., permission errors, filesystem issues, etc.).
// 3. Error Handling for File Creation:
// go
// Copy
// if err != nil {
// 	panic("unable to create file")
// }
// If there was an error in creating the file (err != nil), the program will panic and print the message "unable to create file". A panic is a way to immediately stop program execution, which is typically used for serious errors.
// 4. Deferring the File Closure:
// go
// Copy
// defer f.Close()
// This defer statement ensures that the file will be closed after the function completes (even if there is an error or the function exits early).
// The f.Close() function will be called automatically when the program exits the main function, ensuring proper resource cleanup.
// 5. Creating the Buffer:
// go
// Copy
// buf := make([]byte, 64)
// A buffer (buf) of type []byte (byte slice) is created with a length of 64. This will hold the data read from the file. The buffer can hold up to 64 bytes from the file.
// 6. Reading from the File:
// go
// Copy
// cnt, err := f.Read(buf)
// The f.Read(buf) function is used to read data from the file f into the buf slice.
// It attempts to read up to 64 bytes (the size of buf) from the file.
// The function returns two values:
// cnt: The number of bytes successfully read from the file (can be less than or equal to the size of the buffer, depending on how much data is available in the file).
// err: An error that occurred while reading the file. If the end of the file is reached without any error, err will be nil.
// 7. Error Handling for Reading the File:
// if err != nil {
// 	panic("unable to read file")
// }
// If there was an error while reading the file (err != nil), the program will panic and print the message "unable to read file".
// Note that this code does not check for the "EOF" (End of File) error, which is a normal occurrence when you reach the end of the file. If you want to handle this scenario gracefully, you should check if err is equal to io.EOF.
// 8. Printing the Result:
// go
// Copy
// fmt.Printf("Read %d bytes\n", cnt)
// fmt.Println(string(buf[:cnt]))
// The first fmt.Printf() prints the number of bytes that were successfully read from the file, indicated by cnt.
// The second fmt.Println() converts the byte slice buf into a string using string(buf[:cnt]) and prints the content. Since buf is a byte slice and cnt indicates how many bytes were read, the string(buf[:cnt]) ensures only the valid bytes (those actually read) are converted and printed.