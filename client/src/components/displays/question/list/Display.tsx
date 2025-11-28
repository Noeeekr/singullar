// Components
import Large from "./Large";
import Medium from "./Medium";

// Models
import type { QuestionListDisplayProps } from ".";

export interface QuestionListDisplayyVariants {
    variant?: "lg" | "md"
}

export default ({ 
    variant = "md", 
    ...props 
}: QuestionListDisplayyVariants & QuestionListDisplayProps) => {
    switch(variant) {
        case "lg":
            return <Large {...props} />        
        case "md":
            return <Medium {...props} />
    }
}