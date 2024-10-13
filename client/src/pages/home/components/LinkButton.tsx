import {
    Link,
} from 'react-router-dom'

import { ISideMenuLinkButton } from '../../../types/propsButtons'

import Button from './Button'
import type { ISideMenuButtonProps, IButtonBaseProps } from './Button'

export type ISideMenuLinkProps = ISideMenuLinkButton & IButtonBaseProps;

const LinkButton = (props: ISideMenuLinkProps): JSX.Element => {
    const { href } = props;
     
    const clonedProps: ISideMenuButtonProps = { ...props, type: "button" }
    
    return (
        <Link to={href} style={{ textDecoration: 'none' }}>
            <Button {...clonedProps}/>
        </Link>
    )
}

export default LinkButton