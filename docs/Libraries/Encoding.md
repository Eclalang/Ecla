# Encoding

Encoding library is used to convert data to and from byte-level and textual representation.

## Index

- [AsciiToString(intArr []int) string ](#ascii)
- [DecodeBase64(str string) []int ](#decodebase64)
- [DecodeGob(intArr []int) string ](#decodegob)
- [DecodeHex(str string) []int](#decodehex)
- [EncodeBase64(intArr []int) string](#encodebase64)
- [EncodeGob(str string) []int](#encodegob)
- [EncodeHex(intArr []int) string](#encodehex)
- [StringToAscii(str string) []int](#stringtoascii)

## Ascii
```
function asciiToString(intArr []int) string
```
Converts ascii to string

## DecodeBase64
```
function decodeBase64(str string) []int
```
Decodes base64 to ascii

## DecodeGob
```
function decodeGob(intArr []int) string
```
Decodes gob to string

## DecodeHex
```
function decodeHex(str string) []int
```
Decodes hex to ascii

## EncodeBase64
```
function encodeBase64(intArr []int) string
```
Encodes ascii to base64

## EncodeGob
```
function encodeGob(str string) []int
```
Encodes string to gob

## EncodeHex
```
function encodeHex(intArr []int) string
```
Encodes ascii to hex

## StringToAscii
```
function stringToAscii(str string) []int
```
Converts string to ascii