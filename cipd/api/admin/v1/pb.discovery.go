// Code generated by cproto. DO NOT EDIT.

package api

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"cipd.Admin",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 236, 123, 75, 108, 27, 201,
			182, 152, 250, 67, 138, 46, 255, 164, 242, 79, 166, 199, 242, 185,
			28, 203, 34, 109, 138, 148, 100, 91, 99, 201, 99, 207, 80, 18,
			37, 183, 175, 68, 210, 36, 53, 190, 246, 224, 194, 110, 178, 139,
			100, 141, 201, 46, 222, 238, 166, 100, 141, 227, 204, 234, 5, 200,
			38, 155, 236, 130, 172, 178, 8, 144, 0, 65, 22, 201, 3, 146,
			229, 13, 178, 204, 58, 219, 172, 178, 206, 54, 64, 128, 32, 56,
			213, 93, 77, 202, 159, 153, 185, 147, 32, 171, 49, 198, 152, 62,
			85, 167, 206, 175, 234, 212, 249, 20, 77, 254, 3, 37, 215, 186,
			66, 116, 251, 172, 56, 244, 68, 32, 90, 163, 78, 145, 13, 134,
			193, 113, 65, 130, 244, 124, 56, 89, 80, 147, 153, 105, 146, 40,
			227, 252, 230, 33, 185, 208, 22, 131, 194, 7, 243, 155, 68, 206,
			214, 16, 172, 105, 47, 23, 187, 60, 232, 141, 90, 133, 182, 24,
			20, 187, 162, 111, 187, 221, 49, 155, 97, 112, 60, 100, 126, 200,
			237, 127, 106, 218, 191, 208, 141, 221, 218, 230, 191, 214, 231, 119,
			67, 138, 181, 8, 175, 240, 156, 245, 251, 127, 116, 197, 145, 219,
			68, 252, 167, 255, 123, 134, 36, 169, 57, 63, 117, 119, 134, 252,
			151, 51, 68, 59, 67, 141, 249, 41, 186, 250, 215, 51, 32, 23,
			180, 69, 31, 54, 71, 157, 14, 243, 124, 88, 130, 144, 212, 162,
			15, 142, 29, 216, 192, 221, 128, 121, 237, 158, 237, 118, 25, 116,
			132, 55, 176, 3, 2, 91, 98, 120, 236, 241, 110, 47, 128, 213,
			229, 229, 7, 209, 2, 176, 220, 118, 1, 160, 212, 239, 131, 156,
			243, 193, 99, 62, 243, 14, 153, 83, 32, 208, 11, 130, 161, 191,
			81, 44, 58, 236, 144, 245, 197, 144, 121, 190, 178, 1, 42, 57,
			140, 132, 88, 106, 133, 66, 20, 9, 129, 58, 115, 184, 31, 120,
			188, 53, 10, 184, 112, 193, 118, 29, 24, 249, 12, 184, 11, 190,
			24, 121, 109, 38, 71, 90, 220, 181, 189, 99, 41, 151, 159, 135,
			35, 30, 244, 64, 120, 242, 255, 98, 20, 16, 24, 8, 135, 119,
			120, 219, 70, 10, 121, 176, 61, 6, 67, 230, 13, 120, 16, 48,
			7, 134, 158, 56, 228, 14, 115, 32, 232, 217, 1, 4, 61, 212,
			174, 223, 23, 71, 220, 237, 66, 91, 184, 14, 199, 69, 62, 46,
			34, 48, 96, 193, 6, 33, 128, 127, 110, 127, 32, 152, 15, 162,
			163, 36, 106, 11, 135, 193, 96, 228, 7, 224, 177, 192, 230, 174,
			164, 106, 183, 196, 33, 78, 69, 22, 35, 224, 138, 128, 183, 89,
			30, 130, 30, 247, 161, 207, 253, 0, 41, 76, 114, 116, 157, 15,
			196, 113, 184, 223, 238, 219, 124, 192, 188, 194, 231, 132, 224, 238,
			164, 45, 148, 16, 67, 79, 56, 163, 54, 27, 203, 65, 198, 130,
			252, 95, 201, 65, 32, 210, 206, 17, 237, 209, 128, 185, 129, 173,
			54, 169, 40, 60, 16, 65, 143, 121, 48, 176, 3, 230, 113, 187,
			239, 143, 77, 45, 55, 40, 232, 49, 2, 147, 210, 199, 74, 85,
			24, 151, 43, 145, 176, 107, 15, 24, 10, 52, 121, 182, 92, 49,
			158, 147, 118, 231, 129, 143, 26, 185, 33, 41, 225, 249, 48, 176,
			143, 161, 197, 240, 164, 56, 16, 8, 96, 174, 35, 60, 159, 225,
			161, 24, 122, 98, 32, 2, 6, 161, 77, 2, 31, 28, 230, 241,
			67, 230, 64, 199, 19, 3, 18, 90, 193, 23, 157, 224, 8, 143,
			73, 116, 130, 192, 31, 178, 54, 158, 32, 24, 122, 28, 15, 150,
			135, 103, 199, 13, 79, 145, 239, 75, 217, 9, 52, 159, 88, 13,
			104, 84, 119, 154, 207, 75, 245, 50, 88, 13, 168, 213, 171, 223,
			89, 219, 229, 109, 216, 124, 1, 205, 39, 101, 216, 170, 214, 94,
			212, 173, 221, 39, 77, 120, 82, 221, 219, 46, 215, 27, 80, 170,
			108, 195, 86, 181, 210, 172, 91, 155, 7, 205, 106, 189, 65, 32,
			83, 106, 128, 213, 200, 200, 153, 82, 229, 5, 148, 255, 84, 171,
			151, 27, 13, 168, 214, 193, 218, 175, 237, 89, 229, 109, 120, 94,
			170, 215, 75, 149, 166, 85, 110, 228, 193, 170, 108, 237, 29, 108,
			91, 149, 221, 60, 108, 30, 52, 161, 82, 109, 18, 216, 179, 246,
			173, 102, 121, 27, 154, 213, 188, 100, 251, 241, 58, 168, 238, 192,
			126, 185, 190, 245, 164, 84, 105, 150, 54, 173, 61, 171, 249, 66,
			50, 220, 177, 154, 21, 100, 182, 83, 173, 19, 40, 65, 173, 84,
			111, 90, 91, 7, 123, 165, 58, 212, 14, 234, 181, 106, 163, 12,
			168, 217, 182, 213, 216, 218, 43, 89, 251, 229, 237, 2, 88, 21,
			168, 84, 161, 252, 93, 185, 210, 132, 198, 147, 210, 222, 222, 73,
			69, 9, 84, 159, 87, 202, 117, 148, 126, 82, 77, 216, 44, 195,
			158, 85, 218, 220, 43, 35, 43, 169, 231, 182, 85, 47, 111, 53,
			81, 161, 241, 215, 150, 181, 93, 174, 52, 75, 123, 121, 2, 141,
			90, 121, 203, 42, 237, 229, 161, 252, 167, 242, 126, 109, 175, 84,
			127, 145, 143, 136, 54, 202, 207, 14, 202, 149, 166, 85, 218, 131,
			237, 210, 126, 105, 183, 220, 128, 236, 47, 89, 165, 86, 175, 110,
			29, 212, 203, 251, 40, 117, 117, 7, 26, 7, 155, 141, 166, 213,
			60, 104, 150, 97, 183, 90, 221, 150, 198, 110, 148, 235, 223, 89,
			91, 229, 198, 67, 216, 171, 54, 164, 193, 14, 26, 229, 60, 129,
			237, 82, 179, 36, 89, 215, 234, 213, 29, 171, 217, 120, 136, 223,
			155, 7, 13, 75, 26, 206, 170, 52, 203, 245, 250, 65, 173, 105,
			85, 43, 57, 120, 82, 125, 94, 254, 174, 92, 135, 173, 210, 65,
			163, 188, 45, 45, 92, 173, 160, 182, 120, 86, 202, 213, 250, 11,
			36, 139, 118, 144, 59, 144, 135, 231, 79, 202, 205, 39, 229, 58,
			26, 85, 90, 171, 132, 102, 104, 52, 235, 214, 86, 115, 18, 173,
			90, 135, 102, 181, 222, 36, 19, 122, 66, 165, 188, 187, 103, 237,
			150, 43, 91, 101, 156, 174, 34, 153, 231, 86, 163, 156, 131, 82,
			221, 106, 32, 130, 37, 25, 195, 243, 210, 11, 168, 30, 72, 173,
			113, 163, 14, 26, 101, 18, 126, 79, 28, 221, 188, 220, 79, 176,
			118, 160, 180, 253, 157, 133, 146, 71, 216, 181, 106, 163, 97, 69,
			199, 69, 154, 109, 235, 73, 100, 243, 2, 33, 41, 162, 233, 212,
			128, 212, 21, 252, 74, 81, 35, 51, 245, 144, 156, 34, 122, 106,
			33, 252, 12, 7, 191, 156, 122, 44, 7, 79, 135, 159, 225, 224,
			205, 169, 188, 28, 212, 194, 207, 112, 112, 97, 234, 142, 28, 140,
			62, 195, 193, 91, 83, 25, 57, 72, 194, 207, 112, 112, 113, 234,
			15, 114, 240, 102, 248, 25, 14, 102, 167, 110, 200, 193, 27, 225,
			231, 255, 210, 137, 110, 78, 81, 227, 238, 212, 76, 250, 127, 232,
			80, 130, 46, 115, 153, 199, 219, 32, 227, 39, 12, 152, 239, 219,
			93, 22, 134, 128, 99, 49, 130, 182, 237, 130, 199, 150, 48, 208,
			4, 2, 236, 67, 193, 29, 112, 88, 135, 187, 242, 250, 27, 13,
			251, 24, 76, 152, 67, 78, 174, 151, 215, 239, 177, 24, 121, 80,
			170, 89, 126, 1, 74, 16, 28, 15, 121, 219, 238, 3, 123, 107,
			15, 134, 125, 6, 220, 71, 122, 50, 126, 5, 96, 251, 242, 22,
			243, 216, 95, 70, 204, 15, 8, 68, 183, 154, 199, 252, 161, 112,
			145, 243, 241, 80, 94, 125, 182, 139, 244, 48, 248, 244, 132, 83,
			128, 29, 225, 1, 119, 253, 192, 118, 219, 76, 69, 35, 140, 175,
			188, 205, 96, 71, 8, 120, 23, 14, 1, 120, 195, 54, 108, 218,
			94, 246, 131, 36, 163, 32, 115, 140, 28, 198, 166, 145, 231, 250,
			240, 153, 249, 135, 33, 153, 247, 120, 177, 245, 24, 60, 109, 84,
			43, 50, 146, 48, 63, 190, 230, 59, 194, 131, 215, 18, 251, 53,
			106, 22, 218, 66, 34, 138, 214, 15, 172, 29, 192, 235, 119, 239,
			95, 23, 8, 33, 196, 48, 167, 52, 106, 220, 77, 157, 109, 37,
			37, 155, 187, 228, 223, 22, 201, 141, 15, 83, 167, 128, 15, 152,
			31, 216, 131, 225, 231, 210, 167, 135, 228, 84, 83, 225, 208, 57,
			50, 237, 51, 140, 83, 254, 156, 6, 90, 214, 168, 43, 144, 94,
			36, 9, 215, 118, 133, 63, 167, 131, 150, 77, 212, 67, 96, 243,
			31, 126, 58, 229, 58, 23, 83, 84, 105, 215, 157, 95, 78, 187,
			98, 73, 255, 134, 212, 235, 239, 151, 200, 52, 77, 204, 79, 253,
			35, 77, 251, 61, 247, 250, 61, 247, 250, 61, 247, 250, 61, 247,
			250, 61, 247, 250, 61, 247, 250, 255, 152, 123, 197, 41, 17, 126,
			170, 220, 107, 83, 37, 100, 248, 169, 114, 175, 56, 33, 91, 136,
			19, 178, 91, 83, 69, 149, 144, 225, 167, 202, 189, 226, 132, 108,
			49, 78, 200, 178, 227, 132, 12, 63, 255, 253, 53, 153, 123, 37,
			126, 196, 200, 151, 254, 151, 215, 160, 4, 113, 200, 29, 103, 20,
			62, 216, 48, 20, 220, 13, 228, 173, 198, 7, 24, 101, 28, 54,
			100, 174, 195, 220, 32, 204, 130, 142, 195, 241, 31, 133, 203, 100,
			178, 212, 182, 251, 204, 117, 108, 47, 63, 166, 194, 28, 204, 170,
			162, 60, 64, 222, 158, 29, 207, 110, 143, 99, 132, 154, 192, 16,
			128, 73, 129, 132, 49, 70, 138, 126, 24, 226, 184, 11, 7, 205,
			45, 40, 15, 69, 187, 39, 217, 21, 192, 10, 100, 114, 227, 98,
			100, 193, 248, 135, 183, 176, 188, 63, 107, 158, 232, 179, 97, 192,
			219, 176, 235, 177, 174, 240, 184, 237, 194, 86, 36, 19, 28, 245,
			120, 187, 7, 236, 109, 192, 144, 33, 222, 152, 99, 36, 37, 56,
			129, 150, 221, 126, 115, 100, 123, 142, 76, 11, 143, 153, 237, 129,
			112, 63, 98, 105, 251, 254, 104, 128, 92, 237, 126, 31, 6, 220,
			29, 5, 76, 198, 68, 88, 91, 38, 177, 74, 125, 225, 118, 243,
			192, 11, 172, 0, 125, 102, 15, 199, 170, 122, 12, 50, 254, 128,
			217, 30, 115, 50, 224, 139, 48, 212, 186, 98, 18, 139, 64, 96,
			183, 194, 236, 212, 101, 12, 89, 118, 100, 142, 25, 48, 111, 136,
			81, 84, 6, 8, 168, 203, 244, 131, 251, 209, 101, 189, 188, 188,
			188, 178, 36, 255, 107, 46, 47, 111, 200, 255, 94, 162, 22, 235,
			235, 235, 235, 75, 43, 171, 75, 119, 87, 154, 171, 119, 55, 238,
			175, 111, 220, 95, 47, 172, 171, 63, 47, 11, 4, 54, 143, 209,
			224, 129, 199, 219, 129, 52, 101, 36, 146, 135, 228, 243, 112, 196,
			128, 185, 254, 200, 139, 146, 241, 35, 38, 115, 241, 182, 112, 15,
			153, 23, 64, 32, 72, 180, 171, 98, 0, 80, 223, 217, 130, 187,
			119, 239, 174, 99, 146, 196, 0, 73, 186, 93, 191, 64, 160, 193,
			24, 124, 175, 178, 157, 163, 163, 163, 2, 103, 65, 167, 32, 188,
			110, 209, 235, 180, 241, 47, 46, 42, 4, 111, 131, 63, 103, 127,
			13, 86, 14, 3, 204, 151, 80, 14, 115, 120, 159, 16, 245, 9,
			43, 27, 176, 37, 6, 195, 81, 192, 38, 142, 180, 148, 173, 86,
			109, 88, 127, 130, 215, 120, 130, 178, 57, 204, 129, 101, 120, 29,
			35, 197, 9, 100, 148, 102, 143, 83, 95, 159, 5, 175, 162, 205,
			203, 202, 229, 149, 131, 189, 189, 92, 238, 147, 120, 242, 12, 103,
			151, 115, 15, 39, 100, 90, 253, 37, 153, 186, 44, 64, 42, 162,
			227, 216, 199, 19, 178, 249, 129, 55, 106, 7, 146, 193, 161, 221,
			135, 224, 48, 226, 120, 2, 253, 86, 112, 152, 7, 41, 208, 195,
			223, 170, 210, 97, 33, 56, 68, 232, 231, 52, 10, 145, 70, 62,
			107, 195, 109, 88, 89, 94, 62, 169, 225, 221, 207, 106, 248, 156,
			187, 119, 87, 225, 245, 46, 11, 26, 199, 126, 192, 6, 56, 93,
			242, 119, 120, 159, 53, 79, 110, 196, 142, 181, 87, 110, 90, 251,
			101, 232, 4, 145, 24, 159, 91, 115, 171, 19, 40, 73, 15, 172,
			74, 115, 237, 30, 4, 188, 253, 198, 135, 71, 144, 205, 102, 195,
			145, 92, 39, 40, 56, 71, 79, 120, 183, 183, 109, 7, 114, 85,
			14, 190, 254, 26, 238, 174, 230, 224, 31, 128, 156, 219, 19, 71,
			106, 74, 217, 173, 88, 132, 18, 202, 235, 136, 35, 95, 146, 68,
			207, 90, 89, 94, 158, 184, 151, 252, 66, 140, 192, 228, 125, 180,
			178, 246, 177, 203, 197, 212, 112, 249, 202, 218, 189, 123, 247, 190,
			186, 187, 182, 188, 28, 251, 127, 139, 117, 132, 199, 224, 192, 229,
			111, 21, 149, 245, 175, 150, 63, 164, 82, 248, 109, 155, 153, 13,
			245, 135, 108, 54, 52, 74, 81, 110, 22, 254, 201, 193, 210, 164,
			56, 191, 112, 130, 145, 14, 154, 75, 209, 89, 152, 160, 35, 15,
			64, 238, 196, 1, 184, 247, 217, 3, 240, 212, 62, 180, 225, 117,
			184, 145, 133, 246, 200, 243, 152, 27, 32, 202, 62, 239, 247, 185,
			63, 113, 0, 240, 186, 132, 129, 28, 133, 71, 240, 249, 5, 63,
			115, 204, 225, 209, 120, 180, 224, 178, 163, 205, 17, 239, 59, 204,
			203, 230, 80, 177, 70, 100, 161, 136, 69, 104, 152, 156, 170, 204,
			1, 16, 167, 18, 234, 206, 221, 0, 53, 143, 48, 67, 213, 35,
			181, 165, 5, 114, 133, 22, 82, 150, 178, 140, 109, 112, 255, 179,
			54, 136, 180, 80, 65, 20, 106, 199, 65, 47, 76, 146, 79, 152,
			127, 82, 252, 108, 238, 195, 189, 217, 101, 193, 214, 216, 26, 217,
			156, 188, 1, 101, 105, 191, 111, 15, 135, 220, 237, 18, 2, 150,
			27, 142, 132, 21, 105, 94, 6, 185, 9, 59, 29, 15, 217, 201,
			40, 6, 118, 116, 71, 71, 133, 11, 129, 239, 213, 13, 254, 43,
			47, 226, 136, 85, 1, 154, 24, 27, 184, 159, 15, 201, 132, 163,
			200, 44, 243, 14, 131, 232, 251, 165, 119, 3, 225, 6, 189, 247,
			75, 239, 28, 251, 248, 125, 243, 93, 79, 140, 188, 247, 27, 239,
			6, 220, 125, 191, 241, 206, 103, 237, 247, 223, 23, 222, 97, 98,
			128, 7, 249, 253, 159, 95, 102, 8, 28, 245, 152, 199, 32, 92,
			141, 132, 236, 254, 145, 125, 236, 3, 123, 139, 137, 133, 31, 199,
			253, 142, 24, 121, 224, 240, 46, 15, 124, 140, 240, 125, 6, 17,
			167, 60, 72, 86, 121, 2, 33, 179, 60, 72, 110, 121, 25, 173,
			36, 75, 25, 137, 127, 100, 158, 88, 26, 218, 142, 19, 150, 70,
			193, 145, 80, 212, 152, 221, 238, 21, 100, 167, 69, 101, 44, 118,
			63, 142, 238, 249, 40, 157, 192, 80, 216, 21, 48, 26, 202, 64,
			171, 150, 102, 101, 212, 15, 7, 87, 62, 157, 215, 228, 242, 68,
			242, 23, 195, 144, 114, 200, 41, 243, 50, 3, 254, 168, 211, 225,
			111, 49, 217, 146, 45, 173, 48, 85, 193, 115, 128, 105, 22, 100,
			51, 7, 205, 173, 76, 238, 225, 137, 81, 130, 6, 242, 216, 95,
			70, 220, 99, 78, 1, 74, 16, 182, 116, 194, 195, 224, 203, 122,
			147, 255, 200, 60, 240, 123, 98, 212, 119, 148, 41, 71, 62, 147,
			169, 85, 214, 246, 99, 110, 14, 180, 142, 9, 138, 145, 195, 13,
			112, 177, 194, 115, 131, 40, 191, 250, 240, 40, 161, 33, 237, 19,
			172, 134, 182, 231, 143, 217, 180, 24, 1, 153, 197, 4, 2, 236,
			118, 155, 13, 3, 104, 137, 160, 39, 121, 226, 218, 176, 32, 86,
			58, 248, 31, 201, 1, 182, 11, 162, 211, 241, 89, 24, 239, 119,
			132, 167, 186, 118, 121, 200, 172, 46, 175, 124, 133, 119, 230, 202,
			253, 230, 242, 202, 198, 221, 229, 141, 149, 251, 133, 229, 149, 151,
			153, 232, 116, 251, 32, 225, 248, 210, 29, 218, 126, 64, 64, 98,
			74, 254, 194, 133, 167, 182, 59, 178, 189, 99, 88, 185, 159, 7,
			164, 86, 136, 28, 200, 62, 180, 27, 109, 143, 15, 131, 60, 166,
			126, 39, 146, 29, 27, 48, 104, 168, 94, 154, 204, 147, 48, 251,
			10, 15, 251, 68, 30, 234, 7, 54, 102, 147, 14, 124, 31, 8,
			171, 81, 109, 72, 31, 203, 230, 198, 62, 21, 55, 124, 10, 3,
			241, 35, 239, 247, 109, 233, 92, 204, 93, 58, 104, 20, 29, 209,
			246, 139, 207, 89, 171, 56, 150, 164, 88, 103, 29, 230, 49, 183,
			205, 138, 187, 125, 209, 178, 251, 175, 170, 82, 4, 191, 136, 242,
			20, 39, 152, 252, 153, 196, 93, 73, 75, 93, 52, 121, 233, 230,
			145, 68, 175, 49, 51, 147, 105, 180, 250, 120, 173, 244, 65, 77,
			91, 76, 41, 203, 48, 9, 253, 148, 134, 223, 191, 246, 3, 175,
			35, 87, 78, 40, 36, 218, 126, 97, 24, 222, 107, 168, 202, 106,
			177, 207, 91, 158, 237, 29, 203, 198, 92, 161, 23, 12, 250, 95,
			202, 47, 181, 54, 71, 226, 190, 71, 120, 47, 70, 60, 252, 33,
			107, 195, 226, 194, 139, 165, 133, 193, 210, 130, 211, 92, 120, 178,
			177, 176, 191, 177, 208, 40, 44, 116, 94, 46, 22, 96, 143, 191,
			97, 71, 220, 103, 121, 188, 176, 208, 62, 114, 143, 136, 20, 93,
			182, 134, 123, 12, 158, 10, 199, 150, 71, 117, 209, 135, 239, 95,
			91, 141, 170, 10, 244, 59, 225, 85, 229, 68, 96, 54, 247, 250,
			207, 217, 176, 7, 23, 221, 114, 63, 8, 39, 220, 8, 252, 88,
			66, 169, 138, 246, 144, 203, 253, 80, 163, 82, 157, 98, 40, 107,
			241, 99, 218, 82, 79, 197, 96, 105, 137, 64, 14, 109, 40, 90,
			178, 239, 101, 71, 58, 6, 12, 43, 165, 161, 116, 13, 209, 9,
			27, 223, 118, 232, 100, 202, 193, 252, 240, 66, 142, 77, 47, 91,
			182, 170, 105, 251, 99, 106, 150, 252, 115, 141, 152, 230, 148, 62,
			69, 141, 159, 244, 139, 233, 127, 162, 65, 125, 92, 182, 169, 51,
			47, 58, 242, 168, 75, 235, 250, 220, 109, 79, 230, 28, 228, 211,
			73, 7, 236, 143, 252, 0, 15, 129, 140, 91, 159, 41, 40, 200,
			167, 42, 138, 151, 192, 221, 118, 127, 228, 243, 67, 86, 32, 228,
			44, 73, 160, 116, 38, 53, 127, 210, 127, 188, 64, 206, 132, 96,
			2, 165, 157, 86, 144, 70, 141, 159, 82, 231, 21, 100, 80, 227,
			39, 122, 129, 252, 247, 80, 47, 141, 154, 127, 167, 233, 52, 253,
			95, 53, 168, 8, 119, 201, 101, 93, 59, 224, 135, 236, 100, 237,
			104, 71, 154, 2, 150, 79, 159, 186, 99, 11, 80, 137, 22, 170,
			123, 27, 14, 237, 254, 136, 249, 225, 209, 27, 19, 147, 141, 65,
			63, 224, 253, 62, 244, 236, 67, 6, 238, 36, 79, 73, 58, 90,
			72, 194, 26, 168, 45, 70, 110, 128, 91, 131, 149, 162, 42, 143,
			63, 52, 94, 84, 122, 229, 163, 191, 228, 132, 129, 206, 73, 173,
			53, 147, 38, 254, 78, 211, 127, 186, 24, 25, 76, 75, 72, 189,
			167, 21, 40, 205, 144, 58, 171, 64, 3, 193, 153, 217, 184, 99,
			255, 223, 110, 144, 123, 93, 81, 104, 247, 60, 49, 224, 163, 129,
			60, 186, 253, 81, 155, 23, 237, 225, 144, 185, 93, 238, 178, 226,
			0, 63, 189, 162, 122, 14, 137, 218, 248, 87, 99, 132, 66, 136,
			80, 80, 8, 233, 95, 122, 2, 200, 252, 99, 131, 156, 106, 244,
			108, 207, 177, 220, 142, 160, 23, 73, 130, 187, 14, 123, 43, 27,
			254, 137, 122, 8, 208, 53, 146, 240, 3, 59, 96, 178, 221, 127,
			110, 21, 10, 159, 229, 87, 104, 32, 94, 61, 68, 71, 106, 204,
			243, 132, 55, 103, 128, 150, 61, 85, 15, 1, 122, 143, 76, 183,
			61, 134, 65, 97, 206, 4, 45, 123, 122, 53, 253, 225, 147, 65,
			33, 142, 76, 117, 133, 138, 171, 70, 67, 71, 174, 74, 252, 242,
			170, 8, 149, 230, 137, 193, 2, 123, 46, 249, 139, 43, 16, 141,
			46, 17, 58, 244, 68, 91, 230, 35, 175, 152, 27, 240, 128, 51,
			127, 110, 90, 190, 125, 204, 198, 51, 229, 104, 130, 46, 144, 115,
			129, 8, 236, 254, 24, 53, 37, 81, 207, 202, 209, 24, 45, 75,
			102, 20, 194, 171, 33, 243, 48, 37, 154, 59, 5, 90, 86, 175,
			159, 83, 227, 53, 230, 53, 88, 59, 243, 175, 12, 50, 253, 84,
			180, 228, 78, 156, 35, 58, 119, 162, 119, 23, 157, 59, 191, 121,
			15, 38, 172, 109, 252, 38, 107, 255, 138, 61, 250, 192, 218, 191,
			188, 63, 63, 99, 237, 228, 175, 183, 246, 244, 175, 181, 118, 234,
			83, 214, 166, 95, 147, 164, 143, 7, 223, 159, 187, 8, 70, 246,
			244, 234, 205, 159, 51, 169, 242, 144, 122, 180, 230, 246, 27, 146,
			144, 118, 166, 151, 200, 108, 163, 89, 106, 150, 95, 29, 84, 100,
			31, 119, 199, 42, 111, 207, 76, 209, 51, 36, 213, 104, 150, 234,
			77, 171, 178, 59, 163, 209, 211, 100, 186, 126, 80, 169, 32, 160,
			227, 84, 105, 179, 26, 78, 25, 56, 213, 56, 216, 218, 42, 55,
			26, 51, 38, 77, 17, 115, 167, 100, 237, 205, 36, 112, 88, 34,
			149, 183, 103, 146, 155, 169, 151, 201, 80, 162, 167, 255, 238, 10,
			73, 82, 243, 220, 84, 73, 35, 255, 217, 148, 79, 92, 231, 166,
			232, 234, 127, 52, 79, 188, 86, 173, 60, 144, 217, 231, 222, 193,
			150, 5, 165, 81, 208, 19, 158, 143, 217, 208, 30, 111, 51, 87,
			166, 218, 174, 19, 61, 64, 148, 134, 118, 27, 49, 195, 153, 60,
			124, 199, 60, 159, 11, 23, 86, 11, 203, 144, 69, 132, 76, 52,
			149, 193, 234, 242, 88, 140, 228, 219, 131, 43, 130, 40, 60, 99,
			68, 195, 36, 157, 189, 149, 249, 32, 199, 220, 106, 48, 236, 115,
			27, 195, 83, 156, 16, 68, 52, 10, 4, 94, 68, 20, 226, 16,
			218, 22, 195, 99, 188, 252, 39, 208, 192, 14, 162, 82, 106, 50,
			152, 219, 82, 210, 240, 78, 12, 241, 252, 226, 158, 181, 85, 174,
			52, 202, 75, 171, 133, 101, 66, 224, 192, 237, 51, 127, 156, 55,
			203, 140, 115, 40, 31, 146, 49, 50, 247, 237, 35, 16, 30, 216,
			93, 143, 133, 165, 1, 119, 229, 75, 7, 119, 187, 249, 248, 73,
			100, 226, 201, 230, 132, 153, 148, 100, 220, 63, 129, 32, 31, 131,
			226, 71, 141, 205, 82, 195, 106, 228, 9, 60, 183, 154, 79, 170,
			7, 205, 19, 47, 18, 178, 153, 191, 109, 53, 173, 106, 69, 182,
			219, 75, 149, 23, 240, 71, 171, 178, 157, 135, 232, 53, 40, 170,
			131, 80, 68, 142, 6, 148, 239, 137, 13, 198, 78, 176, 239, 68,
			79, 67, 241, 131, 77, 223, 118, 187, 35, 187, 203, 160, 43, 14,
			153, 39, 223, 206, 199, 175, 54, 178, 61, 75, 160, 207, 7, 60,
			236, 48, 250, 31, 107, 20, 183, 182, 103, 82, 234, 97, 159, 78,
			93, 87, 13, 235, 232, 211, 152, 162, 198, 197, 233, 44, 89, 37,
			122, 98, 138, 154, 115, 83, 160, 165, 111, 129, 60, 253, 97, 220,
			254, 65, 180, 80, 112, 76, 175, 195, 247, 42, 8, 157, 36, 124,
			153, 78, 96, 106, 48, 151, 32, 228, 52, 49, 19, 50, 199, 185,
			170, 207, 97, 158, 144, 8, 179, 134, 171, 250, 5, 5, 233, 212,
			184, 122, 249, 74, 132, 168, 81, 35, 29, 35, 106, 18, 34, 10,
			210, 169, 145, 142, 17, 117, 106, 92, 139, 17, 113, 217, 53, 253,
			148, 130, 112, 46, 70, 52, 168, 241, 69, 140, 104, 104, 8, 41,
			138, 134, 78, 141, 47, 98, 68, 147, 26, 215, 99, 68, 83, 67,
			72, 81, 52, 117, 106, 92, 143, 17, 19, 212, 152, 143, 17, 19,
			26, 66, 73, 5, 233, 212, 152, 143, 17, 147, 212, 184, 17, 35,
			38, 53, 132, 20, 197, 164, 78, 141, 27, 151, 175, 144, 135, 178,
			139, 111, 222, 156, 90, 214, 210, 69, 192, 203, 6, 83, 69, 249,
			226, 216, 146, 207, 116, 34, 76, 252, 186, 125, 22, 218, 55, 54,
			254, 196, 47, 0, 110, 166, 102, 73, 94, 229, 146, 11, 250, 133,
			204, 141, 176, 36, 110, 217, 232, 248, 50, 156, 43, 119, 147, 52,
			38, 115, 187, 5, 253, 230, 100, 110, 183, 112, 34, 183, 91, 72,
			157, 157, 200, 237, 22, 102, 41, 42, 38, 83, 59, 227, 150, 126,
			65, 165, 52, 38, 53, 111, 233, 11, 138, 138, 150, 196, 73, 69,
			5, 183, 240, 86, 76, 69, 51, 168, 113, 107, 150, 146, 45, 73,
			69, 167, 198, 162, 126, 33, 179, 6, 189, 209, 64, 254, 50, 196,
			118, 164, 215, 202, 140, 65, 253, 250, 35, 47, 29, 160, 99, 243,
			62, 115, 162, 19, 6, 194, 237, 31, 43, 29, 116, 147, 154, 139,
			250, 45, 197, 29, 55, 103, 81, 79, 41, 72, 163, 198, 226, 169,
			115, 10, 50, 168, 177, 56, 75, 73, 70, 114, 55, 168, 145, 211,
			179, 153, 75, 97, 141, 204, 3, 56, 178, 125, 136, 66, 161, 34,
			110, 152, 212, 204, 233, 139, 138, 184, 145, 196, 53, 215, 20, 164,
			81, 35, 247, 197, 151, 10, 66, 122, 183, 22, 163, 125, 48, 169,
			113, 91, 207, 102, 110, 156, 32, 30, 69, 76, 232, 219, 126, 216,
			71, 82, 108, 76, 147, 154, 183, 245, 92, 54, 34, 101, 38, 113,
			181, 98, 131, 39, 241, 118, 204, 198, 52, 168, 113, 251, 214, 34,
			89, 147, 108, 18, 212, 184, 163, 103, 85, 157, 207, 3, 232, 112,
			151, 251, 61, 172, 195, 121, 7, 2, 207, 110, 191, 145, 215, 131,
			39, 186, 120, 209, 228, 20, 195, 132, 73, 205, 59, 250, 109, 197,
			48, 145, 68, 58, 138, 33, 158, 232, 59, 95, 220, 80, 144, 65,
			141, 59, 183, 22, 201, 61, 201, 48, 73, 141, 37, 253, 70, 102,
			17, 220, 209, 160, 197, 60, 60, 85, 113, 208, 6, 21, 106, 33,
			232, 141, 124, 232, 216, 158, 98, 151, 52, 169, 185, 164, 223, 81,
			236, 146, 9, 164, 162, 78, 8, 250, 197, 82, 42, 173, 32, 131,
			26, 75, 215, 231, 201, 3, 201, 110, 154, 26, 5, 253, 70, 230,
			14, 200, 112, 63, 193, 52, 102, 37, 60, 88, 90, 1, 222, 129,
			145, 251, 198, 21, 71, 174, 98, 57, 109, 82, 179, 160, 47, 41,
			45, 166, 19, 72, 73, 177, 156, 214, 168, 81, 72, 93, 86, 144,
			65, 141, 194, 245, 249, 200, 164, 41, 106, 20, 245, 27, 153, 28,
			120, 209, 77, 23, 233, 39, 67, 70, 204, 118, 200, 60, 245, 202,
			19, 49, 76, 153, 212, 44, 234, 5, 197, 48, 149, 64, 58, 138,
			97, 74, 163, 70, 49, 53, 167, 32, 131, 26, 197, 235, 243, 228,
			22, 193, 205, 53, 239, 77, 149, 180, 116, 250, 19, 190, 63, 233,
			230, 232, 71, 247, 82, 231, 201, 77, 98, 154, 26, 186, 249, 125,
			157, 102, 174, 192, 200, 229, 127, 25, 49, 121, 23, 115, 7, 165,
			235, 112, 22, 153, 93, 147, 238, 125, 95, 191, 55, 43, 217, 106,
			210, 189, 239, 71, 34, 105, 210, 189, 239, 167, 136, 130, 12, 106,
			220, 159, 153, 37, 139, 146, 188, 70, 141, 53, 157, 102, 210, 128,
			1, 198, 238, 247, 193, 87, 215, 62, 94, 33, 63, 136, 150, 226,
			128, 174, 191, 166, 223, 167, 17, 21, 116, 253, 181, 152, 3, 138,
			188, 22, 185, 190, 38, 93, 127, 109, 102, 86, 58, 159, 134, 174,
			255, 224, 231, 157, 79, 147, 158, 253, 64, 95, 83, 196, 241, 236,
			61, 136, 14, 169, 38, 61, 251, 65, 228, 21, 154, 244, 236, 7,
			145, 243, 105, 8, 172, 255, 90, 231, 211, 164, 143, 175, 235, 15,
			178, 17, 41, 244, 241, 245, 152, 13, 250, 248, 122, 204, 6, 125,
			124, 61, 114, 62, 13, 125, 124, 227, 111, 119, 62, 77, 122, 251,
			134, 190, 174, 24, 162, 183, 111, 196, 12, 209, 219, 55, 34, 231,
			211, 164, 183, 111, 68, 206, 167, 161, 183, 127, 253, 183, 58, 159,
			38, 125, 253, 107, 125, 67, 177, 75, 72, 42, 106, 143, 208, 215,
			191, 142, 156, 79, 147, 190, 254, 117, 228, 124, 26, 218, 251, 209,
			111, 113, 62, 77, 250, 251, 35, 253, 107, 165, 5, 250, 251, 163,
			152, 37, 250, 251, 163, 200, 249, 52, 233, 239, 143, 34, 231, 211,
			208, 223, 31, 255, 237, 206, 167, 73, 111, 127, 172, 63, 82, 12,
			209, 219, 31, 199, 12, 209, 219, 31, 71, 206, 167, 73, 111, 127,
			124, 125, 158, 100, 37, 195, 20, 53, 190, 213, 255, 144, 185, 54,
			62, 225, 120, 220, 127, 16, 173, 69, 149, 206, 40, 15, 73, 153,
			136, 26, 67, 73, 106, 124, 123, 250, 162, 130, 52, 106, 124, 123,
			73, 109, 33, 186, 247, 183, 243, 16, 87, 248, 255, 116, 158, 172,
			126, 178, 194, 111, 243, 161, 83, 180, 135, 188, 104, 59, 3, 238,
			22, 15, 87, 194, 143, 168, 190, 55, 113, 58, 253, 115, 255, 16,
			34, 253, 155, 26, 7, 25, 135, 156, 122, 42, 90, 91, 194, 237,
			240, 46, 189, 73, 204, 55, 220, 13, 107, 206, 115, 171, 51, 5,
			100, 90, 216, 151, 235, 254, 200, 93, 167, 46, 103, 233, 28, 153,
			110, 139, 193, 128, 185, 129, 172, 68, 79, 213, 21, 72, 175, 144,
			105, 199, 59, 126, 229, 141, 92, 89, 105, 166, 234, 73, 199, 59,
			174, 143, 220, 204, 60, 73, 96, 85, 187, 77, 47, 145, 228, 15,
			162, 245, 42, 174, 107, 19, 63, 136, 150, 229, 100, 222, 144, 212,
			83, 209, 10, 171, 169, 69, 146, 108, 75, 113, 36, 202, 233, 213,
			243, 161, 24, 177, 148, 245, 104, 154, 174, 17, 147, 187, 29, 33,
			133, 56, 189, 154, 249, 153, 218, 45, 170, 168, 235, 18, 63, 243,
			111, 52, 114, 166, 105, 119, 119, 248, 219, 58, 27, 10, 47, 160,
			121, 146, 232, 240, 183, 12, 101, 194, 42, 240, 114, 200, 112, 18,
			5, 129, 122, 136, 148, 22, 196, 104, 218, 93, 58, 67, 140, 225,
			155, 80, 198, 83, 117, 252, 164, 105, 146, 82, 191, 25, 141, 12,
			19, 195, 244, 58, 33, 45, 79, 188, 97, 238, 171, 192, 238, 70,
			205, 144, 83, 225, 8, 18, 187, 70, 78, 73, 226, 114, 214, 12,
			215, 202, 129, 166, 221, 189, 253, 23, 66, 198, 123, 64, 175, 145,
			43, 251, 165, 90, 173, 92, 127, 133, 213, 195, 7, 37, 231, 101,
			66, 203, 149, 131, 253, 114, 29, 171, 209, 90, 105, 235, 143, 165,
			221, 114, 99, 70, 163, 87, 200, 133, 29, 196, 222, 47, 237, 237,
			84, 235, 251, 229, 237, 87, 205, 210, 110, 99, 70, 199, 210, 181,
			252, 167, 90, 181, 222, 148, 3, 175, 154, 213, 87, 155, 207, 102,
			140, 213, 191, 106, 36, 81, 194, 227, 71, 115, 228, 212, 158, 61,
			114, 219, 189, 167, 162, 69, 63, 220, 138, 244, 233, 120, 192, 218,
			166, 69, 146, 42, 181, 132, 23, 32, 230, 228, 68, 250, 242, 71,
			125, 0, 249, 131, 86, 122, 155, 156, 222, 101, 65, 188, 241, 39,
			214, 156, 139, 129, 112, 114, 153, 156, 221, 225, 111, 247, 109, 239,
			141, 52, 138, 127, 18, 155, 126, 188, 101, 155, 137, 151, 134, 61,
			228, 79, 255, 89, 84, 46, 31, 252, 94, 46, 255, 94, 46, 255,
			63, 45, 151, 207, 196, 229, 114, 122, 92, 46, 167, 199, 229, 242,
			77, 249, 169, 81, 227, 210, 244, 35, 242, 159, 52, 162, 39, 167,
			168, 249, 197, 212, 45, 45, 253, 247, 26, 72, 239, 66, 203, 132,
			45, 233, 82, 205, 10, 127, 204, 217, 58, 134, 45, 171, 182, 29,
			255, 218, 220, 30, 227, 9, 207, 31, 255, 104, 200, 197, 32, 204,
			152, 172, 176, 161, 180, 181, 23, 110, 209, 73, 236, 168, 127, 141,
			135, 79, 253, 86, 180, 117, 12, 204, 117, 160, 221, 231, 204, 13,
			124, 249, 68, 233, 177, 69, 31, 92, 17, 255, 78, 138, 200, 19,
			105, 7, 188, 197, 251, 60, 56, 150, 63, 40, 229, 62, 139, 74,
			249, 36, 166, 138, 95, 164, 206, 146, 38, 49, 147, 178, 196, 156,
			215, 239, 164, 119, 33, 188, 36, 152, 15, 54, 12, 194, 215, 107,
			153, 135, 202, 230, 58, 123, 107, 15, 184, 203, 124, 245, 187, 89,
			188, 217, 194, 95, 93, 251, 129, 240, 88, 28, 203, 11, 50, 160,
			38, 195, 106, 115, 62, 121, 94, 65, 88, 70, 207, 92, 85, 144,
			65, 141, 249, 155, 57, 114, 87, 242, 215, 168, 1, 250, 90, 250,
			22, 88, 46, 15, 184, 124, 92, 181, 101, 198, 236, 133, 63, 104,
			155, 20, 38, 38, 143, 185, 40, 36, 207, 41, 72, 167, 6, 156,
			191, 168, 32, 131, 26, 112, 227, 30, 41, 72, 242, 58, 53, 50,
			122, 62, 253, 7, 168, 71, 63, 228, 31, 103, 6, 159, 164, 140,
			242, 100, 146, 179, 10, 194, 229, 244, 138, 130, 12, 106, 100, 50,
			183, 35, 195, 97, 5, 173, 175, 166, 119, 97, 71, 26, 35, 27,
			222, 68, 220, 245, 185, 19, 30, 225, 158, 237, 58, 125, 230, 229,
			32, 176, 187, 62, 12, 228, 197, 135, 219, 135, 115, 93, 126, 200,
			92, 8, 163, 220, 9, 254, 120, 218, 22, 146, 23, 20, 164, 83,
			99, 225, 162, 50, 28, 230, 171, 11, 55, 151, 201, 66, 216, 195,
			185, 61, 181, 170, 165, 175, 66, 99, 52, 196, 219, 146, 57, 147,
			250, 76, 182, 109, 110, 39, 206, 143, 219, 54, 119, 244, 249, 137,
			182, 205, 29, 253, 234, 68, 219, 230, 206, 23, 215, 201, 119, 170,
			109, 179, 164, 95, 77, 91, 176, 61, 26, 12, 199, 191, 98, 198,
			100, 106, 104, 183, 223, 200, 127, 197, 17, 8, 216, 45, 149, 161,
			47, 186, 126, 94, 190, 74, 50, 63, 56, 33, 2, 116, 60, 123,
			192, 142, 132, 247, 166, 64, 38, 58, 64, 75, 250, 197, 137, 14,
			208, 210, 149, 57, 178, 167, 58, 64, 69, 61, 157, 254, 6, 118,
			184, 235, 132, 54, 147, 103, 207, 17, 238, 98, 0, 67, 219, 247,
			225, 59, 187, 207, 49, 237, 183, 162, 160, 220, 180, 187, 242, 213,
			26, 109, 43, 223, 228, 7, 49, 39, 84, 161, 168, 95, 154, 104,
			33, 21, 231, 174, 146, 29, 213, 66, 90, 209, 231, 210, 235, 80,
			126, 139, 150, 243, 165, 94, 146, 33, 119, 209, 5, 97, 147, 119,
			159, 141, 152, 119, 28, 253, 226, 111, 145, 189, 13, 45, 140, 193,
			221, 95, 140, 121, 224, 94, 173, 196, 141, 47, 220, 171, 149, 203,
			87, 200, 221, 176, 5, 180, 54, 181, 161, 165, 23, 97, 155, 117,
			164, 227, 28, 161, 42, 39, 61, 43, 122, 146, 119, 196, 68, 235,
			103, 45, 53, 27, 245, 100, 166, 168, 241, 149, 126, 121, 162, 179,
			243, 149, 190, 22, 119, 118, 146, 56, 121, 102, 162, 179, 243, 213,
			217, 217, 137, 206, 206, 87, 23, 47, 225, 241, 15, 59, 59, 15,
			244, 75, 153, 63, 128, 237, 181, 120, 224, 217, 222, 241, 135, 221,
			153, 240, 199, 38, 100, 162, 249, 243, 64, 255, 74, 213, 210, 90,
			2, 215, 167, 38, 154, 63, 15, 78, 205, 76, 52, 127, 30, 92,
			184, 24, 137, 171, 99, 153, 117, 97, 162, 137, 179, 174, 63, 184,
			52, 209, 196, 89, 15, 123, 106, 81, 19, 103, 125, 122, 178, 137,
			179, 62, 75, 201, 90, 88, 60, 63, 158, 250, 86, 75, 223, 6,
			75, 21, 191, 242, 50, 80, 57, 216, 167, 188, 86, 21, 211, 143,
			83, 103, 165, 36, 178, 152, 254, 38, 146, 36, 172, 153, 191, 209,
			31, 159, 159, 168, 153, 191, 57, 81, 51, 127, 147, 58, 55, 81,
			51, 127, 51, 75, 201, 109, 162, 155, 58, 53, 183, 167, 44, 45,
			61, 15, 219, 44, 176, 121, 223, 143, 75, 248, 143, 185, 163, 62,
			219, 169, 25, 178, 79, 76, 83, 71, 238, 59, 250, 149, 244, 183,
			80, 245, 120, 151, 227, 29, 143, 91, 29, 230, 187, 121, 188, 71,
			219, 65, 255, 24, 108, 95, 213, 177, 254, 168, 21, 253, 51, 143,
			64, 64, 156, 167, 69, 207, 181, 186, 148, 127, 71, 223, 14, 139,
			102, 93, 110, 252, 142, 126, 90, 65, 26, 53, 118, 206, 196, 115,
			6, 53, 118, 46, 93, 38, 15, 165, 28, 26, 53, 158, 232, 75,
			233, 2, 68, 63, 80, 250, 168, 230, 15, 127, 42, 210, 239, 159,
			108, 245, 134, 92, 241, 20, 60, 209, 119, 174, 68, 148, 181, 36,
			18, 251, 82, 65, 72, 250, 102, 86, 65, 6, 53, 158, 220, 201,
			147, 2, 193, 234, 214, 220, 159, 58, 208, 210, 25, 168, 51, 127,
			212, 151, 247, 183, 55, 114, 101, 140, 62, 145, 250, 69, 150, 67,
			255, 217, 79, 225, 109, 96, 154, 134, 49, 69, 205, 138, 222, 48,
			36, 93, 195, 64, 221, 42, 228, 44, 57, 75, 146, 8, 161, 93,
			171, 230, 5, 114, 158, 76, 135, 160, 73, 205, 170, 89, 57, 79,
			206, 169, 129, 4, 34, 144, 49, 172, 81, 163, 122, 250, 220, 24,
			54, 168, 81, 157, 165, 49, 61, 141, 26, 53, 115, 46, 166, 135,
			26, 215, 204, 234, 133, 24, 31, 79, 126, 109, 130, 30, 106, 93,
			59, 61, 49, 111, 80, 163, 118, 249, 74, 76, 79, 167, 198, 51,
			51, 29, 211, 67, 15, 120, 102, 214, 230, 98, 124, 244, 129, 103,
			19, 244, 80, 128, 103, 167, 47, 141, 97, 131, 26, 207, 230, 174,
			146, 108, 68, 207, 160, 70, 221, 188, 154, 185, 138, 217, 81, 38,
			131, 69, 121, 116, 96, 28, 214, 103, 178, 171, 162, 88, 25, 38,
			53, 235, 230, 179, 116, 76, 202, 72, 224, 218, 49, 43, 52, 115,
			253, 244, 197, 49, 140, 180, 175, 204, 73, 119, 49, 208, 176, 205,
			48, 10, 32, 96, 34, 68, 20, 148, 164, 70, 243, 244, 57, 5,
			105, 212, 104, 134, 81, 22, 33, 131, 26, 205, 43, 115, 170, 42,
			254, 63, 1, 0, 0, 255, 255, 26, 109, 134, 118, 229, 63, 0,
			0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("cipd.Admin")
	if err != nil {
		panic(err)
	}
	return ret
}
