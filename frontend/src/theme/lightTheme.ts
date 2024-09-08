import {
    createTheme,
} from '@mui/material'


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
    typography: {
        h4: {
            fontSize: 21,
            fontFamily: 'Roboto',
            fontWeight: 'bold'
        },
        body2: {
            fontSize: 14.5,
        }
    },
    palette: {
        primary: {
            lightPurple: 'rgb(154, 61, 230)',
            darkPurple: 'rgb(114, 41, 230)',
            semiLight: 'rgb(150,150,150)',
            lightGray: 'rgb(100,100,100)',
            contrast: 'rgb(255, 102, 0)',
            light: 'rgb(255,255,255)',
            dark: 'rgb(0,0,0)',
            main: 'rgb(200,0,200)', // not defined, it is just necessary to eixt
        }
    },
    components: {
        MuiFormControl: {
            styleOverrides: {
                root: {
                    width: '100%' // general, not theme specific
                }
            }
        },
        MuiOutlinedInput: {
            styleOverrides: {
                root: {
                    backgroundColor: 'rgb(250,250,255)',
                    height: '43px', // general, not theme specific
                    borderRadius: '10px', // general, not theme specific
                    '&:hover .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'rgb(215,215,215)',
                    },
                    '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'rgb(156,90,220)',
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
                    color: 'rgb(120,120,120)',
                    translate: '0px -6px',
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