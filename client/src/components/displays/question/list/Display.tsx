// Components
import Large from "./Large";
import Medium from "./Medium";

// Models
import type { QuestionListDisplayProps } from ".";

export interface QuestionListDisplayVariants {
    variant?: "lg" | "md"
}

export default ({ 
    variant = "md", 
    ...props 
}: QuestionListDisplayVariants & QuestionListDisplayProps) => {
    switch(variant) {
        case "lg":
            return <Large {...props} />        
        case "md":
            return <Medium {...props} />
    }
}