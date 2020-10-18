# afmt

(c) 2020 Gon Yi  
<https://gonyyi.com/copyright.txt>
 

## Intro

afmt is a library for string formats with a goal of zero allocation.


---


## Extract

Assume there is a data such as `https://blog.gonyyi.com/myLife/123/`, 
and you want to extract `myLife` and `123` from it. This can be easily done 
using an `Extract` function that utilize a regular expression.

Syntax:

> `afmt.NewExtract`(`regex`, `numOfRef`)


```go
// Create an extract function `extFn` with the regular expression.
// And we are expecting 2 references from it. The function created here can be run
// over and over without recompile.
extFn := afmt.NewExtract("^https://blog.gonyyi.com/([^/]+)/([^/]+)/?$", 2)

// This will actually use the function created from `NewExtract` to extract necessary
// data. If `2` for _number of expected field_ is given when creating a function above,
// this will return a slice size of 2 _regardless of it matches or not_. This is to prevent
// unexpected out of index error.
out1 := extFn("https://blog.gonyyi.com/myLife/123/") // returns string[]{"myLife", "123"}

// When no match is found, it will still return an empty string slice with expected size.
out2 := extFn("https://gonyyi.com/myLife/123/") // returns string[]{"", ""}
```

Also there is a simpler `Extract` function, but if any of this needs to be repeated, this can waste resources
as it will need to compile over and over.
_(To save resources, use `NewExtract()`` to create an object and reuse its regex)_

- `Extract`(s `string`, fmtr `string`, numFlds `int`) `[]string`
    - Arguments
        - s: input string
        - fmtr: format regex
        - numFlds: number of extracted fields
    - Returns: array of string

NewExtract() returns a function that has already compiled regular expression. This is suitable when repeated


---

## Humanize

Filesize (in byte) to string
(eg. 1024*1024 --> 1M)

```go
// afmt.HumanBytes(BYTES_IN_INT64, NUM_DECIMAL)
afmt.HumanBytes(1024*1024, 1) // returns "1.0M"
```

Humanize number
(eg. 1000 --> 1K)

```go
// afmt.HumanNumber(NUM_INT64, NUM_DECIMAL)
afmt.HumanNumber(10000, 1) // returns "10.0K"
afmt.HumanNumber(100350000000, 1) // returns "100.3B"
afmt.HumanNumber(100350000001, 1) // returns "100.4B" -- rounded after exact half.
```

TestNumberWithComma
(eg. 1000000 --> 1,000,000)

```go
// afmt.NumberWithComma(NUM_INT64)
afmt.NumberWithComma(1000000) // returns "1,000,000"
```


---

## Shorter String

Sometimes, especially when creating a log, there can be too much information. Let's say,
you are writing a verbose log for your program. And you like to dump a RSA certificate
to the screen just to make sure for each records you have different public keys. As the
certificate is too long, it will take multiple lines making it hard for you to compare.
So you make peak little bit such as first 50 bytes. However, if the certificate wasn't
in the file, then you can have an error if you haven't considered the case. This is
to save something like that. Example usage includes a long business name or an address,
aka. large single string:

`123 Main St, Conway, AR 72034` into `123 Main...72034`


Usage:

> afmt.NewShorterFunc( `Target Output Length`, `Between Marker`, 
> `Min Length of Left Side`, `Min Length of Right Side` )


Example:

- Input is `123456789ABCDEF` _(total of 16 bytes)_
- But you want it to be:
    - total of `9 bytes`
    - delimiter _(between marker)_ to be `..`
    - min left side to be 4 bytes,
    - min right side to be 2 bytes.
- Function would be:
    ```go
    shorter := afmt.NewShorterFunc(9, "..", 4, 2)
    shorter("123456789ABCDEF") // returns "12345..EF"


For different minimum left and right side length, see the table below:
_(assume input is `123456789ABCDEF` and the marker is `..`)_

| Target Length | MinLeft | MinRight | Output       |
|:--------------|:--------|:---------|:-------------|
| 9             | 4       | 2        | `12345..EF`  |
| 9             | 2       | 4        | `12..BCDEF`  |
| 9             | 4       | 4        | `1234..DEF`  |
| 10            | 4       | 4        | `1234..CDEF` |
| 10            | 0       | 0        | `1234..CDEF` |
| 10            | 0       | 1        | `..89ABCDEF` |
| 10            | 1       | 0        | `12345678..` |
| 10            | 3       | 20       | `123..BCDEF` |



---

## Formatted String

`Hyphenate` is a function that will format the string with a delimiter.

> afmt.Hyphenate( `Raw Data String`, `Delimiter`, `Size`... )

This function returns `string` and `bool` -- if the returned boolean is `false`,
it means length is unexpected.

---

Let's say you are working on a SSN. And your data has it without any delimiter/separator.
And you need to format a string like "123456789" into "123-45-6789":

> ```go
> afmt.Hyphenate("123456789", "-", 3, 2, 4) // returns `123-45-6789`
> ```

If you are working on a VISA card number, it will be total of 16 digits separated by every
4th with a hyphen.

> ```go
> afmt.Hyphenate( "1234123412341234", "-", 4, 4, 4, 4 ) // returns "1234-1234-1234-1234"
> ```

__Note:__ the function will return a `false` for its `bool` output IF unexpected length
of string is given.
