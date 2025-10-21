// Icons;
import { BiDirections } from 'react-icons/bi';
import { LuScanFace, LuAtom } from 'react-icons/lu';
import { MdOutlineNotificationsNone } from 'react-icons/md';
import { FaChalkboardTeacher, FaEdit } from 'react-icons/fa';
import { IoSchool, IoPeopleSharp, IoExitOutline } from 'react-icons/io5';
import { GoPerson, GoLock, GoDiscussionDuplicate } from 'react-icons/go';
import { TbSmartHome, TbSpeakerphone, TbGridDots } from 'react-icons/tb';

// components;
import NotificationPopupContent from '../../../popups/NotificationPopupContent'
import MyAccountPopupContent from '../../../popups/PopupContentMyAccount'
import QuickAccessPopupContent from '../../../popups/PopupContentQuickAccess'

// Data
import {
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
    sidemenu_data_utilities,
} from './any_role_routes'
import { ButtonProps } from '@components/buttons/Button/Button';
import { LinkButtonProps } from '@components/buttons/Link';

import { SideMenuButtons } from '../SideMenuButtons';
import { NavigationPopupProps } from '@components/layout/navbar/AppNavbar';
/**
 * This must reflect all the id's fields present in navbar pop-up data 
 **/
export type AdminNavbarPopupIds = "notifications" | "schooldata" | "myaccount" | "" | "help";

export interface INavbarPopupData {
    variant: "side" | "bubble",
    title: string,
    id: AdminNavbarPopupIds,
    icon: JSX.Element,
    element: JSX.Element,
};

const admin_sidemenu_data_myaccount: (ButtonProps | LinkButtonProps)[] = [
    {
        title: "Dados pessoais e acesso",
        icon: {
            component: <LuScanFace />
        },
        type: "link",
        href: "/admin/notdone",
    },
    {
        title: "Comunicações",
        icon: {
            component: <TbSpeakerphone />,
        },
        type: "link",
        href: "/admin/notdone",
    },
    {
        title: "Privacidade",
        icon: {
            component: <GoLock />
        },
        type: "link",
        href: "/admin/notdone",
    },
    {
        title: "Sair",
        icon: {
            component: <IoExitOutline />
        },
    },
];

export const admin_navbar_popup_data: NavigationPopupProps[] = [
    {
        id: "notifications",
        title: "Notificações",
        variant: "side",
        icon: {
            component: <MdOutlineNotificationsNone color="white" fontSize={22} />
        },
        element: <NotificationPopupContent />,
        type: "popup",
    },
    {
        id: "myaccount",
        title: "Minha conta",
        variant: "bubble",
        icon: {
            component: <GoPerson color="white" fontSize={22} />,
        },
        element: <MyAccountPopupContent menus={admin_sidemenu_data_myaccount} />,
        type: "popup",
    },
    {
        id: "help",
        variant: "side",
        title: "Central de ajuda",
        element: <div>Lorem Ipsum</div>,
        icon: {
            component: <BiDirections color="white" fontSize={20} />,
        },
        type: "popup",
    },
    {
        id: "quickaccess",
        title: "Acesso rápido",
        variant: "bubble",
        icon: {
            component: <TbGridDots color="white" fontSize={23} />,
        },
        element: <QuickAccessPopupContent menus={sidemenu_data_quickaccess.menus} />,
        type: "popup",
    },
];

export const admin_sidemenu_data: SideMenuButtons[] = [
    {
        title: "Principal",
        menus: [
            {
                title: "Inicio",
                type: "link",
                href: "/home",
                icon: {
                    component: <TbSmartHome />,
                },
            },
            ...admin_navbar_popup_data,
        ],
        type: "group",
    },
    {
        title: "Ferramentas",
        menus: [
            {
                title: "Dados escolares",
                type: "link",
                icon: {
                    component: <LuAtom color="white" fontSize={22} />,
                },
                href: "/admin/dashboard",
            },
            {
                title: "Minhas turmas",
                type: "link",
                href: "/admin/classes",
                icon: {
                    component: <IoSchool />,
                },
            },
            {
                title: "Meus professores",
                type: "link",
                href: "/admin/teachers",
                icon: {
                    component: <FaChalkboardTeacher />,
                },
            },
            {
                title: "Meus alunos",
                type: "link",
                href: "/admin/students",
                icon: {
                    component: <IoPeopleSharp />,
                },
            },
        ],
        type: "group",
    },
    {
        title: "Plataforma",
        menus: [
            {
                title: "Banco de questões",
                type: "link",
                href: "/platform/questions",
                icon: {
                    component: <FaEdit />,
                },
            },
            {
                title: "Estudo orientado",
                type: "link",
                href: "/platform/study",
                icon: {
                    component: <GoDiscussionDuplicate />,
                },
            },
            ...sidemenu_data_classroom.menus,
        ],
        type: "group",
    },
    sidemenu_data_utilities,
];

