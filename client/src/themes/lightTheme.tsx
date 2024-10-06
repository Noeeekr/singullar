import { createTheme } from '@mui/material/styles'

import './fonts.css'

declare module '@mui/material/styles' {
    interface PaletteColor { // Extra Types definitions for typescript
        purpleExtraLight?: string,
        purpleLightInv?: string,
        purpleLight?: string,
        purpleDark?: string,

        paperLight?: string,
        
        whiteNone?: string,
        whiteLow?:  string,
        whiteSemiLow?: string,
        whiteMedium?:  string,
        whiteSemiMedium?: string,
        whiteHigh?:  string,

        contrast?:  string,
    }
    interface SimplePaletteColorOptions { // Extra types config
        purpleExtraLight?: string,
        purpleLightInv?: string
        purpleLight?: string,
        purpleDark?: string,

        paperLight?: string,
        
        whiteNone?: string,
        whiteLow?:  string,
        whiteSemiLow?: string,
        whiteMedium?:  string,
        whiteSemiMedium?: string,
        whiteHigh?:  string,

        contrast?:  string,
    }
    interface BreakpointOverrides {
        xs: true;
        sm: true;
        md: true;
        lg: true;
        xl: true;
    }
}

const lightPaletteTheme = createTheme({
    palette: {
        // todo: primary: purple colors
        // todo: secondary: undefined for now
        // todo: error: red colors
        // todo: contrast: orange, black
        primary: {
            main: 'rgb(255,255,255)', // not part of theme : necessary value

            purpleExtraLight: 'rgba(170,148,240)',
            purpleLightInv: 'rgba(220,210,240,0.4)',
            purpleLight: 'rgb(154, 61, 230)',
            purpleDark: 'rgb(114, 41, 230)',

            paperLight: 'rgb(245, 245, 250)',

            whiteNone: 'rgb(0,0,0)',
            whiteLow: 'rgb(100,100,100)',
            whiteSemiLow: 'rgb(120,120,120)',
            whiteMedium: 'rgb(150,150,150)',
            whiteSemiMedium: 'rgb(150,150,150)',
            whiteHigh: 'rgb(240,240,240)',

            contrast: 'rgb(255, 102, 0)',
        },
        error: {
            main: 'rgb(255, 77, 106)',
            light: 'rgb(255, 227, 232)',
            dark: 'rgb(212, 7, 40)',
        }
    },
})

const lightTheme = createTheme({
    breakpoints: {
        values: {
            xs: 760,          // Extra-small devices 
            sm: 900,          // Small devices 
            md: 1200,         // Medium devices 
            lg: 1365,         // Large devices 
            xl: 1536,         // Extra-large devices 
        }
    },
    palette: lightPaletteTheme.palette,
    components: {
        MuiSvgIcon: {
            styleOverrides: {
              root: {
                fontSize: '1.6rem', // Set a default size
                WebkitFontSmoothing: 'antialiased',
                MozOsxFontSmoothing: 'grayscale',
              },
            },
        },
        MuiTypography: {
            styleOverrides: {
                root: {
                    color: 'rgb(38, 41, 48)',
                    fontFamily: 'inter, system-ui',
                },
                body1: {
                    fontSize: 12
                },
                subtitle1: {
                    fontSize: 12,
                    fontWeight: 600,
                },
                subtitle2: {
                    fontSize: 14,
                    fontWeight: 'bold',
                },
                body2: {
                    fontSize: 14
                },
                h6: {
                    fontSize: 16
                },
                h5: {
                    fontSize: 18
                },
                h4: {// This now will be the older h6
                    fontSize: 20
                },
                h3: { 
                    fontSize: 22
                },
                h2: {
                    fontSize: 26
                },
                h1: {
                    fontSize: 30
                }
            }
        },
        MuiFormControl: {
            styleOverrides: {
                root: {
                    width: '100%' // applies for all themes
                }
            }
        },
        MuiInputBase: {
            styleOverrides: {
                root: {
                    justifyContent: 'end',
                    '& input:-webkit-autofill': {
                        position: 'absolute',
                        top: 0,
                        left: 0,
                        height: '10px',
                        width: 'calc(100% - 16px)',
                    },
                }
            }
        },
        MuiOutlinedInput: {
            styleOverrides: {
                root: {
                    zIndex: 2,
                    height: '43px', // applies for all themes
                    borderRadius: '10px', // applies for all themes
                    backgroundColor: 'rgb(250,250,255)',
                    '&:hover .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'rgb(215,215,215)',
                    },
                    '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'rgb(156,90,220)',
                    },
                    '&.Mui-error .MuiOutlinedInput-notchedOutline': {
                        borderColor: lightPaletteTheme.palette?.error?.main,
                        borderWidth: '2px !important',
                    },
                    '&.Mui-focused.Mui-error .MuiOutlinedInput-notchedOutline': {
                        borderWidth: '2px !important',
                    },
                    '&.Mui-error:hover .MuiOutlinedInput-notchedOutline': {
                        borderColor: lightPaletteTheme.palette?.error?.main,
                    },
                    '&.Mui-disabled': {
                        backgroundColor: 'rgb(243,243,243)'
                    },
                    '&.Mui-disabled:hover .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'rgba(0, 0, 0, 0.26)',
                    }
                },
                notchedOutline: {
                    borderColor: 'rgb(230,230,230)',
                },
            }
        },
        MuiInputLabel: {
            styleOverrides: {
                outlined: {
                    zIndex: 4,
                    translate: '0px -6px',
                    color: 'rgb(120,120,120)',
                    transition: 'linear 150ms all',
                    '&.MuiInputLabel-shrink': {
                        translate: '0px 1px',
                    },
                    '&.Mui-focused': {
                        translate: '0px 1px',

                        color: 'rgb(120,120,120)',
                    }
                },

            },
        },
    }
})

export default lightTheme