package base64zero

import (
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		source   []byte
		expected []byte
	}{
		{
			source:   []byte("A"),
			expected: []byte("QQ==")},
		{
			source:   []byte("Hello, World!"),
			expected: []byte("SGVsbG8sIFdvcmxkIQ==")},
		{
			source:   []byte("mIx3d T3xT"),
			expected: []byte("bUl4M2QgVDN4VA==")},
		{
			source:   []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."),
			expected: []byte("TG9yZW0gSXBzdW0gaXMgc2ltcGx5IGR1bW15IHRleHQgb2YgdGhlIHByaW50aW5nIGFuZCB0eXBlc2V0dGluZyBpbmR1c3RyeS4gTG9yZW0gSXBzdW0gaGFzIGJlZW4gdGhlIGluZHVzdHJ5J3Mgc3RhbmRhcmQgZHVtbXkgdGV4dCBldmVyIHNpbmNlIHRoZSAxNTAwcywgd2hlbiBhbiB1bmtub3duIHByaW50ZXIgdG9vayBhIGdhbGxleSBvZiB0eXBlIGFuZCBzY3JhbWJsZWQgaXQgdG8gbWFrZSBhIHR5cGUgc3BlY2ltZW4gYm9vay4gSXQgaGFzIHN1cnZpdmVkIG5vdCBvbmx5IGZpdmUgY2VudHVyaWVzLCBidXQgYWxzbyB0aGUgbGVhcCBpbnRvIGVsZWN0cm9uaWMgdHlwZXNldHRpbmcsIHJlbWFpbmluZyBlc3NlbnRpYWxseSB1bmNoYW5nZWQuIEl0IHdhcyBwb3B1bGFyaXNlZCBpbiB0aGUgMTk2MHMgd2l0aCB0aGUgcmVsZWFzZSBvZiBMZXRyYXNldCBzaGVldHMgY29udGFpbmluZyBMb3JlbSBJcHN1bSBwYXNzYWdlcywgYW5kIG1vcmUgcmVjZW50bHkgd2l0aCBkZXNrdG9wIHB1Ymxpc2hpbmcgc29mdHdhcmUgbGlrZSBBbGR1cyBQYWdlTWFrZXIgaW5jbHVkaW5nIHZlcnNpb25zIG9mIExvcmVtIElwc3VtLg=="),
		},
		{
			source:   []byte("375949384751294752347623"),
			expected: []byte("Mzc1OTQ5Mzg0NzUxMjk0NzUyMzQ3NjIz"),
		},
		{
			source:   []byte("^$*^%*#&^$&*^#*^%#)()*)(*+_+"),
			expected: []byte("XiQqXiUqIyZeJCYqXiMqXiUjKSgpKikoKitfKw=="),
		},
	}

	for _, tc := range testCases {
		result := Encode(tc.source)
		if string(result) != string(tc.expected) {
			t.Errorf("Expected %s, but got %s", tc.expected, result)
		}
	}
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		source      []byte
		expected    []byte
		expectedErr error
	}{
		{
			source:   []byte("QQ=="),
			expected: []byte("A"),
		},
		{
			source:   []byte("SGVsbG8sIFdvcmxkIQ=="),
			expected: []byte("Hello, World!"),
		},
		{
			source:   []byte("bUl4M2QgVDN4VA=="),
			expected: []byte("mIx3d T3xT"),
		},
		{
			source:   []byte("TG9yZW0gSXBzdW0gaXMgc2ltcGx5IGR1bW15IHRleHQgb2YgdGhlIHByaW50aW5nIGFuZCB0eXBlc2V0dGluZyBpbmR1c3RyeS4gTG9yZW0gSXBzdW0gaGFzIGJlZW4gdGhlIGluZHVzdHJ5J3Mgc3RhbmRhcmQgZHVtbXkgdGV4dCBldmVyIHNpbmNlIHRoZSAxNTAwcywgd2hlbiBhbiB1bmtub3duIHByaW50ZXIgdG9vayBhIGdhbGxleSBvZiB0eXBlIGFuZCBzY3JhbWJsZWQgaXQgdG8gbWFrZSBhIHR5cGUgc3BlY2ltZW4gYm9vay4gSXQgaGFzIHN1cnZpdmVkIG5vdCBvbmx5IGZpdmUgY2VudHVyaWVzLCBidXQgYWxzbyB0aGUgbGVhcCBpbnRvIGVsZWN0cm9uaWMgdHlwZXNldHRpbmcsIHJlbWFpbmluZyBlc3NlbnRpYWxseSB1bmNoYW5nZWQuIEl0IHdhcyBwb3B1bGFyaXNlZCBpbiB0aGUgMTk2MHMgd2l0aCB0aGUgcmVsZWFzZSBvZiBMZXRyYXNldCBzaGVldHMgY29udGFpbmluZyBMb3JlbSBJcHN1bSBwYXNzYWdlcywgYW5kIG1vcmUgcmVjZW50bHkgd2l0aCBkZXNrdG9wIHB1Ymxpc2hpbmcgc29mdHdhcmUgbGlrZSBBbGR1cyBQYWdlTWFrZXIgaW5jbHVkaW5nIHZlcnNpb25zIG9mIExvcmVtIElwc3VtLg=="),
			expected: []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."),
		},
		{
			source:   []byte("Mzc1OTQ5Mzg0NzUxMjk0NzUyMzQ3NjIz"),
			expected: []byte("375949384751294752347623"),
		},
		{
			source:   []byte("XiQqXiUqIyZeJCYqXiMqXiUjKSgpKikoKitfKw=="),
			expected: []byte("^$*^%*#&^$&*^#*^%#)()*)(*+_+"),
		},
		{
			source:      []byte("VGhpc0'lzVGV4dA=="),
			expectedErr: ErrInvalidBase64,
		},
	}

	for _, tc := range testCases {
		result, err := Decode(tc.source)

		if err != nil && err != tc.expectedErr {
			t.Errorf("Unexpected error: %v", err)
		}

		if string(result) != string(tc.expected) {
			t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
		}
	}
}
