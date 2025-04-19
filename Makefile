sample: _build
	${LCC} samples/main.c samples/basic_tile_data.c -o _build/sample.gb

_build:
	mkdir _build || echo "already exists"