// Features
import { useAppSelector } from '../slices/store'
import { useMemo } from 'react'

// Components
import {
    Stack,
    Typography,
    Grid2 as Grid,
} from '@mui/material'
import BannerSlider from './BannerSlider';
import LinkButton from './ButtonLink';

// Data
import { admin_sidemenu_data } from '../pages/admin/data'               
import { supervisor_sidemenu_data } from '../pages/supervisor/data'
import { teacher_sidemenu_data } from '../pages/teacher/data'
import { students_sidemenu_data } from '../pages/student/data'

// Types
import type { 
    ISideMenuItems,
    ISideMenuLinkButton,
} from '../types/buttonProps';

const Root = (): JSX.Element => {
    const { user, mostVisitedUrls } = useAppSelector((state) => state.user);

    const data = useMemo(() => {
        let allItems: ISideMenuItems[];
        let linkItems: ISideMenuLinkButton[] = [];
        let mostVisited: ISideMenuLinkButton[] = [];

        switch (user?.role) {
            case "student":
                allItems = students_sidemenu_data;
                break;
            case "admin":
                allItems = admin_sidemenu_data;
                break;
            case "supervisor":
                allItems = supervisor_sidemenu_data;
                break;
            case "teacher":
                allItems = teacher_sidemenu_data;
                break;
            default:
                allItems = [];
        }

        // Returns only the items;
        allItems.forEach(i => i.items.forEach(z => {
            if (z.type === "link") linkItems.push(z as ISideMenuLinkButton); 
        }))
        
        mostVisited = linkItems.filter(i => mostVisitedUrls[i.href]);
        // Doesnt select the higest on count must become a reducer

        return mostVisited.length ? mostVisited.slice(0,4) : linkItems.slice(0,4)
    }, [user?.role])

    return (
        <Stack component="div" gap={4} paddingBottom={3}>
            <Stack gap={1}>
                <Typography variant="h3" fontWeight="700" my={0}>
                    Olá, administrador
                </Typography>
                <Typography variant="subtitle2" fontWeight="500" color="primary.whiteSemiHigh">
                    {user ? user?.email : ""}
                </Typography>
            </Stack>
            <BannerSlider />
            {
                data && 
                <Typography variant="h6" fontWeight="600" sx={{ marginTop: 8 }}>
                    Mais acessados
                </Typography>
            }
            <Grid container spacing={3}>
                {
                    data.map((d,i) => (
                        <Grid key={i} size={{ mobile: 6, xs: 4, md: 3 }}>
                            <LinkButton
                                showIcon={true}
                                href={d.href}
                                type={d.type}
                                title={d.title}
                                icon={d.icon}
                                variant="paper"
                            />
                        </Grid>
                    ))
                }
                
            </Grid>
        </Stack>
    )
}

export default Root;
