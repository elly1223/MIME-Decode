package main

import "strings"

// Decode decodes an RFC 2047 encoded-word. : 디코딩은 RFC 2047 인코딩된 워드를 디코딩한다.
func (d *WordDecoder) Decode(word string) (string, error) {
	// HasPrefix : 문자열이 접두사인지 판단
	// HasSuffix : 문자열이 접미사인지 판단
	// Count : 문자열이 몇 번 나오는지 구함
	if len(word) < 8 || !strings.HasPrefix(word, "=?") || !strings.HasSuffix(word, "?=") || strings.Count(word, "?") != 4 {
		return "", errInvalidWord
	}

	// 1.
	word = word[2 : len(word)-2]

	// 2.
	// split delimits the first 2 fields : 처음 두 필드를 구분
	// IndexByte : byte 자료형으로 위치를 구함
	split := strings.IndexByte(word, '?')

	// 3.
	// split word "UTF-8?q?ascii" into "UTF-8", 'q', and "ascii"
	// 단어 "UTF-8?q?asci"를 "UTF-8", 'q' 및 "asci"로 분할
	charset := word[:split] //
	if len(charset) == 0 {
		return "", errInvalidWord
	}
	if len(word) < split+3 {
		return "", errInvalidWord
	}

	// 4.
	encoding := word[split+1]
	// the field after split must only be one byte
	// 분할 후 필드는 1byte여야 한다.
	if word[split+2] != '?' {
		return "", errInvalidWord
	}

	// 5.
	text := word[split+3:]

	content, err := decode(encoding, text)
	if err != nil {
		return "", err
	}

	var buf strings.Builder

	if err := d.convert(&buf, charset, content); err != nil {
		return "", err
	}

	// 6.
	return buf.String(), nil
}
