# Fuzzing in Go

Fuzzing is a type of automated testing which continuously manipulates inputs to a program to find bugs. 

Go fuzzing uses coverage guidance to intelligently walk through the code being fuzzed to find and report failures to the user. 

Since it can reach edge cases which humans often miss, fuzz testing can be particularly valuable for finding security exploits and vulnerabilities.

[More info](https://go.dev/doc/security/fuzz/) about fuzzing in go.

### Quick start

1. Write a fuzz-test. Name convention is FuzzXXX and signature `(f *testing.F)`
2. Provide a seed corpus (appropriate values for function):
~~~
testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) 
	}
~~~

3. Finally, write a Fuzz function, which takes a simple arguments like: int, string, byte and etc.
~~~
f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
~~~

4. Run the test:
~~~
go test -fuzz=Fuzz
~~~

5. If the test failed, try to find more info at `reverse/fuzz/FuzzReverse/a265dd251690629f` (example)
6. Fix the test and run again the certain case:
~~~
go test -run=FuzzReverse/a265dd251690629f
~~~
7. Keep in mind that you can limit the Fuzzing work like this:
~~~
go test -fuzz=Fuzz -fuzztime 30s
~~~

Keep in mind, that not every function you can Fuzz, but it's better to try, than not!