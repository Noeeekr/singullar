import {
    MouseEventHandler,
} from 'react'

import {
    Box,
    Stack,
    styled,
    Divider,
    Typography,
} from '@mui/material'

import { GrBook } from "react-icons/gr";
import { FaChalkboardTeacher } from "react-icons/fa";
import { FaRegPenToSquare } from "react-icons/fa6";
import { IoNewspaperOutline } from "react-icons/io5";
import { AiOutlineQuestionCircle } from "react-icons/ai";
import { TbSmartHome } from "react-icons/tb";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { BiDirections } from "react-icons/bi";
import { LuPartyPopper } from "react-icons/lu";
import { FaGithub } from "react-icons/fa";
import { FaShareAlt } from "react-icons/fa";

import MenuItems from './MenuItems'
import { ISideMenuItems } from '../../../types/sideMenu'

// Seriously MUI, what is this syntax???
const SectionTitle = styled(({ children, ...other }: { children: string }) => (
    <Typography variant={"subtitle1"} {...other}>{children}</Typography>
))(({ theme }) => ({
    margin: '0px 10px',
    textTransform: 'uppercase',
    textWrap: 'nowrap',
    color: theme.palette.primary.whiteSemiLow,
}));

const BorderIcon = styled(({ children, style = {}, ...other }: { children: JSX.Element, style?: object }) => (
    <Box style={{ ...style, color: 'black' }} {...other}>
        {
            children
                ? children
                : <Box sx={{
                    borderRadius: 20,
                    backgroundColor: "rgb(110,110,110)",
                    width: 28,
                    height: 28,
                }} />
        }
    </Box>
))(() => ({
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    width: 44,
    height: 44,
    borderRadius: 8,
    border: 'solid 1px gray',
    overflow: 'hidden',
}))

const menuItems_Main: ISideMenuItems = {
    title: "",
    items: [
        {
            title: "Início",
            icon: <TbSmartHome />,
            type: "link",
            href: "/",
        },
        {
            title: "Notificações",
            icon: <MdOutlineNotificationsNone />, // might need the other version for hover effect
            type: "link",
            href: "/",
        },
        {
            title: "Ajuda",
            icon: <BiDirections />,
            type: "link",
            href: "/",
        },
        {
            title: "Minha conta",
            icon: <IoNewspaperOutline />,
            type: "group",
            items: [
                {
                    title: "Dados pessoais e acesso",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Responsáveis vinculados",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Código de acesso",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Comunicações",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Privacidade",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Sair",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
            ],
        },
    ],
}
const menuItems_Classroom: ISideMenuItems = {
    title: "Sala de aula",
    items: [
        {
            title: "Biblioteca de conteúdos",
            icon: <GrBook />,
            type: "link",
            href: "/",
        },
        {
            title: "Atividades",
            icon: <FaRegPenToSquare />, // might need the other version for hover effect
            type: "link",
            href: "/",
        },
        {
            title: "Aulas digitais",
            icon: <FaChalkboardTeacher />,
            type: "link",
            href: "/",
        },
        {
            title: "Simulados e Provas",
            icon: <IoNewspaperOutline />,
            type: "group",
            items: [
                {
                    title: "Avaliações",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Resultados de Avaliações",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
            ],
        },
        {
            title: "Dúvidas e materiais",
            icon: <AiOutlineQuestionCircle />,
            type: "group",
            items: [
                {
                    title: "Ver materiais e tirar dúvidas",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Minhas dúvidas",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
            ],
        },
    ],
}
const menuItems_QuickAccess: ISideMenuItems = {
    title: "Acesso Rápido",
    items: [
        {
            title: "Ir para o perfil do criador",
            icon: <BorderIcon><FaGithub /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 1",
            icon: <BorderIcon><LuPartyPopper /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 2",
            icon: <BorderIcon><FaShareAlt /></BorderIcon>,
            type: "link",
            href: "/",
        },
    ]
}

/*
* onHover must be a toggle type of effect to work.
*/
const SideMenu = (
    { isMobile, isOpen, onHoverOpen }: { isMobile?: boolean, isOpen?: boolean, onHoverOpen?: Function }
): JSX.Element => {

    return (
        <Box
            component="div"
            sx={{
                overflow: 'hidden',
                transition: isMobile ? 'none' : 'width 300ms ease-in-out',
                backgroundColor: 'white',
                borderRight: (theme) => `2px ${theme.palette.primary.whiteHigh} solid`
            }}
            width={
                isMobile
                    ? '100vw'
                    : isOpen
                        ? 250
                        : 55
            }
            padding={isMobile ? 2 : 1 }
            onMouseEnter={isOpen ? undefined : onHoverOpen as MouseEventHandler<HTMLDivElement>}
            onMouseLeave={onHoverOpen as (MouseEventHandler<HTMLDivElement> | undefined)}            
        >
            <div style={{ overflow: 'hidden' }}>
                <Stack
                    component="nav"

                    gap={3}
                >
                    { /* INICIO */}

                    <Box
                        sx={{
                            marginTop: isMobile ? 0 : 3,
                        }}
                    >
                        {
                            <MenuItems
                                showIcon={true}
                                isCompacted={!isOpen}
                                items={
                                    isMobile
                                        ? menuItems_Main.items
                                        : [menuItems_Main.items[0]]
                                }
                            />
                        }
                    </Box>
                    {
                        isOpen && !isMobile 
                        ? <></>
                        : <Divider sx={{
                            marginX: 1,
                            borderBottomWidth: 2,
                            borderColor: 'rgb(240,240,240)',
                        }} />
                    }

                    { /* SALA DE AULA */}

                    <Stack gap={1}>
                        {
                            isOpen && <SectionTitle>
                                {menuItems_Classroom.title}
                            </SectionTitle>
                        }
                        <MenuItems
                            showIcon={true}
                            isCompacted={!isOpen}
                            items={menuItems_Classroom.items}
                        />
                    </Stack>
                    {/* ACESSO RÁPIDO */}
                    {
                        isMobile && (
                            <>
                                <Divider sx={{
                                    borderBottomWidth: 2,
                                    borderColor: 'rgb(240,240,240)'
                                }} />
                                <Stack gap={1}>
                                    <SectionTitle>
                                        {menuItems_QuickAccess.title}
                                    </SectionTitle>
                                    <MenuItems
                                        showIcon={true}
                                        isCompacted={!isOpen}
                                        items={menuItems_QuickAccess.items}
                                    />
                                </Stack>
                            </>
                        )
                    }
                </Stack>
            </div>

        </Box>
    )
}

export default SideMenu