#include "gb/gb.h"

extern unsigned char background_data [];
extern unsigned int background_data_size;
extern unsigned char background_tiles [];
extern unsigned int background_tiles_width;
extern unsigned int background_tiles_height;


int main() {
    set_bkg_data(0, background_data_size, background_data);
    set_bkg_tiles(0, 0, background_tiles_width, background_tiles_height, background_tiles);
    SHOW_BKG;
    return 0;
}
