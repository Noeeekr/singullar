import { createTheme } from '@mui/material/styles'


declare module '@mui/material/styles' {
    interface PaletteColor {
        lightPurple?: string,
        darkPurple?: string,
        lightGray?: string,
        semiLight?: string,
        contrast?: string,
    }
    interface SimplePaletteColorOptions {
        lightPurple?: string,
        darkPurple?: string,
        lightGray?: string,
        semiLight?: string,
        contrast?: string,
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
        primary: { // purple, white, black
            main: 'rgb(255,255,255)', // not defined, it is just necessary to eixt
            
            lightPurple: 'rgb(154, 61, 230)',
            darkPurple: 'rgb(114, 41, 230)',
            
            semiLight: 'rgb(150,150,150)',
            light: 'rgb(255,255,255)',
            
            dark: 'rgb(0,0,0)',
            lightGray: 'rgb(100,100,100)',
            
            contrast: 'rgb(255, 102, 0)',
        },
        error: { // red
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
        MuiTypography: {
            styleOverrides: {
                body1: {
                    fontSize: 14
                },
                body2: {
                    fontSize: 12
                },
                h6: {
                    fontSize: 20
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