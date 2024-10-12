import { GrBook } from "react-icons/gr";
import { FaChalkboardTeacher, FaGithub, FaShareAlt } from "react-icons/fa";
import { FaRegPenToSquare } from "react-icons/fa6";
import { IoNewspaperOutline, IoPeopleOutline, IoExitOutline, IoBarcodeOutline } from "react-icons/io5";
import { AiOutlineQuestionCircle } from "react-icons/ai";
import { TbSmartHome } from "react-icons/tb";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { BiDirections } from "react-icons/bi";
import { LuPartyPopper, LuScanFace } from "react-icons/lu";
import { GoPerson, GoLock } from "react-icons/go";
import { TbSpeakerphone } from "react-icons/tb";

import { 
    Box,
    styled,
} from '@mui/material'

import { ISideMenuItems, ISideMenuLinkButton } from '../../types/propsButtons'

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

export const sideMenuLinks_MyAccount: ISideMenuLinkButton[] = [
    {
        title: "Dados pessoais e acesso",
        icon: <LuScanFace/>,
        type: "link",
        href: "/",
    },
    {
        title: "Responsáveis vinculados",
        icon: <IoPeopleOutline/>,
        type: "link",
        href: "/",
    },
    {
        title: "Código de acesso",
        icon: <IoBarcodeOutline/>,
        type: "link",
        href: "/",
    },
    {
        title: "Comunicações",
        icon: <TbSpeakerphone/>,
        type: "link",
        href: "/",
    },
    {
        title: "Privacidade",
        icon: <GoLock/>,
        type: "link",
        href: "/",
    },
    {
        title: "Sair",
        icon: <IoExitOutline/>,
        type: "link",
        href: "/",
    },
];

export const menuItems_Main: ISideMenuItems = {
    title: "",
    items: [
        {
            title: "Início",
            icon: <TbSmartHome />,
            type: "link",
            href: "/home",
        },
        {
            title: "Notificações",
            icon: <MdOutlineNotificationsNone />, // might need the other version for hover effect
            type: "popup",
            content: <div>DIVINISSIMA</div>,
            notifications: true,
        },
        {
            title: "Ajuda",
            icon: <BiDirections />,
            type: "popup",
            content: <div>DIVINISSIMA</div>,
        },
        {
            title: "Minha conta",
            icon: <GoPerson />,
            type: "group",
            items: sideMenuLinks_MyAccount,
        },
    ],
};
export const menuItems_Classroom: ISideMenuItems = {
    title: "Sala de aula",
    items: [
        {
            title: "Biblioteca de conteúdos",
            icon: <GrBook />,
            type: "link",
            href: "/home/bookshelf",
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
};
export const menuItems_QuickAccess: { title: string, items: ISideMenuLinkButton[] } = {
    title: "Acesso Rápido",
    items: [
        {
            title: "Ir para o perfil do criador",
            description: "Aproveite para ver outros projetos!",
            
            icon: <BorderIcon><FaGithub /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 1",
            description: "Caso não tenha achado o que procurava.",
            icon: <BorderIcon><LuPartyPopper /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 2",
            description: "A sorte vem na terceira tentativa.",

            icon: <BorderIcon><FaShareAlt /></BorderIcon>,
            type: "link",
            href: "/",
        },
    ]
};