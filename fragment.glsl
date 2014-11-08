#version 330

in vec2 texCoord;

layout(location = 0) out vec4 outColor;

uniform sampler2D img;

void main() {
	outColor = texture(img, texCoord);
}

