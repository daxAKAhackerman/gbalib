# --- Project details -------------------------------------------------

PROJ    := pageflip
TARGET  := $(PROJ)

# --- Build defines ---------------------------------------------------

TINYGO_BUILD  := tinygo build

TYNYGO_TARGET   := -target=gameboy-advance


.PHONY : build clean

# --- Build -----------------------------------------------------------
build: $(TARGET).gba

$(TARGET).gba : %.gba : %.go
	$(TINYGO_BUILD) $(TYNYGO_TARGET) -o $@ .
		
# --- Clean -----------------------------------------------------------

clean : 
	@rm -fv *.gba

#EOF
