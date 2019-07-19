// Package dsaext provides datastructures and algorithms
//
// Standard comparison functions that implement the dsaext compareFn
// interface for various data types
//
// @version 2018-07-25
// @author  Robert Altnoeder (r.altnoeder@gmx.net)
//
// Copyright (C) 2018 Robert ALTNOEDER
//
// Redistribution and use in source and binary forms,
// with or without modification, are permitted provided that
// the following conditions are met:
//
//  1. Redistributions of source code must retain the above copyright notice,
//	 this list of conditions and the following disclaimer.
//  2. Redistributions in binary form must reproduce the above copyright
//	 notice, this list of conditions and the following disclaimer in
//	 the documentation and/or other materials provided with the distribution.
//  3. The name of the author may not be used to endorse or promote products
//	 derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
// TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
// EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package dsaext

func CompareInt(value1st, value2nd interface{}) int {
	num1st := value1st.(int)
	num2nd := value2nd.(int)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareUInt8(value1st, value2nd interface{}) int {
	num1st := value1st.(uint8)
	num2nd := value2nd.(uint8)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareUInt16(value1st, value2nd interface{}) int {
	num1st := value1st.(uint16)
	num2nd := value2nd.(uint16)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareUInt32(value1st, value2nd interface{}) int {
	num1st := value1st.(uint32)
	num2nd := value2nd.(uint32)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}


func CompareUInt64(value1st, value2nd interface{}) int {
	num1st := value1st.(uint64)
	num2nd := value2nd.(uint64)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareInt8(value1st, value2nd interface{}) int {
	num1st := value1st.(int8)
	num2nd := value2nd.(int8)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}


func CompareInt16(value1st, value2nd interface{}) int {
	num1st := value1st.(int16)
	num2nd := value2nd.(int16)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}


func CompareInt32(value1st, value2nd interface{}) int {
	num1st := value1st.(int32)
	num2nd := value2nd.(int32)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}


func CompareInt64(value1st, value2nd interface{}) int {
	num1st := value1st.(int64)
	num2nd := value2nd.(int64)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareFloat32(value1st, value2nd interface{}) int {
	num1st := value1st.(float32)
	num2nd := value2nd.(float32)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}


func CompareFloat64(value1st, value2nd interface{}) int {
	num1st := value1st.(float64)
	num2nd := value2nd.(float64)
	var result int = 0
	if num1st < num2nd {
		result = -1
	} else if num1st > num2nd {
		result = 1
	}
	return result;
}

func CompareString(value1st, value2nd interface{}) int {
	text1st := value1st.(string)
	text2nd := value2nd.(string)
	var result int = 0
	if text1st < text2nd {
		result = -1
	} else if text1st > text2nd {
		result = 1
	}
	return result;
}
