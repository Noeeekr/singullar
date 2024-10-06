import {
    useState
} from 'react'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery
} from '@mui/material'
import {
    useTheme
} from '@mui/material/styles'

import NavbarItemPopup from './PopupIconButton'
import UserProfile from './UserProfile'
import Divider from './Divider'
import MenuIcon from './NavbarIconButton'
import SideMenuItems from './SideMenuItems'
import SectionTitle from './SectionTitle'
import { ISideMenuLinkProps } from './LinkButton'

import { IoIosHelpCircleOutline, IoIosNotificationsOutline } from "react-icons/io";
import { TbGridDots } from "react-icons/tb";
import { VscAccount } from "react-icons/vsc";

import {
    sideMenuLinks_MyAccount,
    menuItems_QuickAccess,
} from '../data'


interface IAppNavBarProps {
    menuButtonCallback: Function,
    children?: JSX.Element[],
    showMenu?: boolean
}

const MyAccountPopupContent = (props: { items: ISideMenuLinkProps[] }) => {
    const { items } = props;

    return (
        <Stack>
            <UserProfile structure="center" />
            <Divider sx={{ marginX: 2 }} />
            <Box padding={1}>
                <SideMenuItems
                    gap={1}
                    fontWeight={400}
                    fontSize={13.5}
                    showIcon={true}
                    hasHoverEffect={true}
                    items={items.slice(0, -1)}
                />
            </Box>
            <Divider />
            <Box padding={1}>
                <SideMenuItems
                    gap={1}
                    fontWeight={400}
                    fontSize={13.5}
                    showIcon={true}
                    hasHoverEffect={true}
                    items={[items[items.length - 1]]}
                />
            </Box>
        </Stack>
    )
}
const QuickAccessPopupContent = (props: { items: ISideMenuLinkProps[] }) => {
    const { items } = props;

    return (
        <Box
            sx={{
                backgroundColor: 'rgb(220,220,220,0.2)',
                borderRadius: 4,
                padding: 2,
            }}
        >
            <Stack gap={1}>
                <SectionTitle>
                    Ir para
                </SectionTitle>
                <SideMenuItems
                    showIcon={true}
                    showDescription={true}
                    items={items}
                />
            </Stack>
        </Box>
    )
}
/**
 * Children are shown when navbar menu button is clicked.
 */
const AppNavbar = (props: IAppNavBarProps): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const [isOpen, setIsOpen] = useState<"quickaccess" | "main" | "myaccount" | "help" | "">("")

    const toggleIsOpen = (key: "quickaccess" | "main" | "myaccount" | "help" | "") => {
        if (key === isOpen) {
            setIsOpen("")
        } else {
            setIsOpen(key)
        }
    }

    const {
        children = [],
        menuButtonCallback,
        showMenu,
    } = props;


    return (
        <Box
            sx={{
                zIndex: 4,
            }}
        >
            <Box
                sx={{
                    top: '0',
                    left: '0',

                    width: '100vw',
                    height: isMobile ? '100vh' : 'auto',
                    overflow: isMobile ? 'auto' : 'visible',

                }}
            >
                <Box
                    display="flex"
                    alignItems="center"
                    padding={1}
                    gap={1}
                    height={55}
                    sx={{
                        backgroundColor: (theme) => theme.palette.primary.purpleDark
                    }}
                >
                    <MenuIcon
                        onClick={() => { menuButtonCallback() }}
                        notifications={true}
                    />
                    <Typography
                        component="h3"
                        variant="h5"
                        color="primary.light"
                        sx={{
                            fontSize: 22,
                            fontWeight: 'bold',
                        }}
                    >
                        Singullar
                    </Typography>

                    { /* Desktop */}
                    {
                        !isMobile &&
                        <Stack
                            direction="row"
                            alignItems="center"
                            justifyContent="center"
                            gap={1}
                            margin="auto 0 auto auto"
                        >
                            <NavbarItemPopup
                                structure="side"
                                title="Central de ajuda"
                                id="help"
                                isOpen={isOpen}
                                onClickCb={toggleIsOpen}
                                icon={<IoIosHelpCircleOutline color="white" fontSize={26} />}
                            >
                                <div
                                >Ddd</div>
                            </NavbarItemPopup>
                            <NavbarItemPopup
                                structure="side"
                                title="Notificações"
                                id="notifications"
                                isOpen={isOpen}
                                onClickCb={toggleIsOpen}
                                icon={<IoIosNotificationsOutline color="white" fontSize={26} />}
                            >
                                <div
                                >Ddd</div>
                            </NavbarItemPopup>
                            <NavbarItemPopup
                                title="Minha conta"
                                id="myaccount"
                                isOpen={isOpen}
                                onClickCb={toggleIsOpen}
                                icon={<VscAccount color="white" fontSize={22} />}
                            >
                                <MyAccountPopupContent items={sideMenuLinks_MyAccount} />
                            </NavbarItemPopup>
                            <NavbarItemPopup
                                title="Acesso Rápido"
                                id="quickaccess"
                                isOpen={isOpen}
                                onClickCb={toggleIsOpen}
                                icon={<TbGridDots color="white" fontSize={23} />}
                            >
                                <QuickAccessPopupContent items={menuItems_QuickAccess.items} />
                            </NavbarItemPopup>
                        </Stack>
                    }
                </Box>

                { /* Mobile */}

                <Box>
                    {
                        isMobile && showMenu
                            ? (children.map((child) => {
                                return child
                            }))
                            : <></>
                    }
                </Box>
            </Box>
        </Box>
    )
}

export default AppNavbar