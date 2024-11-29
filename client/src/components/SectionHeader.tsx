// Components
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import CircularButton from '@components/ButtonCircular';
import { Link } from 'react-router-dom'
// Features
import { useNavigate, useLocation } from 'react-router-dom';
import { useMemo } from 'react';
        
/**
 * Contains the title, a button to navigate back and a text subtitle that is a breadcrumbs by default.
 * 
 * Loads children in right side.
 */
const SectionHeader = ({ children, title, subtitle }: { children?: JSX.Element, title: string, subtitle?: string }): JSX.Element => {
    const navigate = useNavigate();
    const pathname = useLocation().pathname;

    const breadcrumbs: string[] = useMemo(() => {
        let pathArray = pathname.slice(1).split("/");

        // Checks if there's a number instead of text in breadcrumbs
        for (let i = 0; i < pathArray.length; i++) {
            if (!Number.isNaN(Number(pathArray[i]))) {
                return pathArray.slice(0,i)
            }
        };
        
        return pathArray;
    },[])

    return (
        <Stack direction="row" alignItems="start" width="100%" gap={1}>
            <CircularButton
                onClickCb={() => { navigate(-1) }}
            />
            <Stack gap={1.5}>
                <Typography component="h4" variant="h3" fontWeight="600">
                    { title }
                </Typography>
                <Typography component="p" variant="body1" fontWeight="600" color="primary.whiteLow">
                    {   
                        subtitle || breadcrumbs.map((urlPart, i) => (
                            <Link 
                                key={urlPart + urlPart}
                                to={"/" + urlPart}
                                style={{
                                    textDecoration: "none",
                                }}
                            >   
                                <Typography sx={{
                                    display: "inline",

                                    color: "primary.whiteLow",
                                    cursor: 'pointer',
                                    
                                    '&:hover': {
                                        color: "primary.purpleLight",
                                    },
                                }}>
                                    { urlPart } { i != breadcrumbs.length - 1 ? "> " : "" }
                                </Typography>
                            </Link>
                        ))
                    }
                </Typography>
            </Stack>
            {
                children || <></>
            }
        </Stack>
    )
}

export default SectionHeader;