import { styled, Box } from "@mui/material"
import React, { ChangeEventHandler, DetailedHTMLProps, FieldsetHTMLAttributes, FocusEventHandler, forwardRef, LabelHTMLAttributes, MutableRefObject, Ref, RefCallback, TextareaHTMLAttributes, useRef, useState } from "react"
import { UseFormRegisterReturn } from "react-hook-form"

export interface FieldsetProps extends FieldsetHTMLAttributes<HTMLFieldSetElement> {
    isFocused?: boolean
    isEmpty?: boolean
    height?: number
}

export interface LabelProps extends LabelHTMLAttributes<HTMLLabelElement> {
    isFocused?: boolean
    isEmpty?: boolean
}

const FieldSet = styled("fieldset")<FieldsetProps>(({
    theme,
    isEmpty,
    isFocused,
    height
}) => ({
    position: "relative",

    backgroundColor: "rgb(250, 250, 255)",
    width: "100%",
    height: height || 100,
    paddingLeft: isFocused || isEmpty ? "13px" : 1,
    paddingRight: isFocused || isEmpty ? "13px" : 1,
    paddingTop: isFocused || isEmpty ? "12px" : 0,
    paddingBottom: isFocused || isEmpty ? "12px" : 0,
    border: isFocused ? `solid 2px ${theme.palette.primary.purpleDark}` : isEmpty ? "solid 1px rgb(180,180,180,0.27)" : "none",
    borderRadius: 10,
    boxShadow: "none",
    outline: "none",

    transition: "background-color 200ms ease-in-out",
    boxSizing: "border-box",
}))

const Label = styled("label")<LabelProps>(({
    isFocused,
    isEmpty,
}) => ({
    position: "absolute",
    top: isFocused || isEmpty ? -8 : 12,
    left: isFocused || isEmpty ? 20 : 14,
    zIndex: 10,
    fontSize: isFocused || isEmpty ? 11 : 16,
    transition: "all 0.2s ease-in-out",
    color: "grey",
    selectable: "none",
    pointerEvents: "none",
}))

export interface TextAreaProps extends DetailedHTMLProps<TextareaHTMLAttributes<HTMLTextAreaElement>, HTMLTextAreaElement> {
    isEmpty?: boolean
    isFocused?: boolean
}

const TextArea = styled("textarea")<TextAreaProps>(({
    isEmpty,
    isFocused,
}) => ({
    position: "absolute",
    top: 0,
    left: 0,

    fontFamily: "Arial, sans-serif",
    fontSize: 16,
    resize: "none",

    backgroundColor: isFocused || isEmpty ? "transparent" : "rgb(250, 250, 255)",
    width: "100%",
    boxShadow: "0px 0px 0px 0px transparent",
    height: "",
    paddingLeft: isFocused ? "13px" : "14px",
    paddingRight: isFocused ? "13px" : "13px",
    paddingTop: isFocused ? "12px" : "13px",
    paddingBottom: isFocused ? "12px" : "12px",
    border: isFocused || isEmpty ? "none" : "solid 2px rgb(180,180,180,0.27)",
    outline: "none",

    boxSizing: "border-box",
    borderRadius: 10,

    overflow: "hidden",

    margin: 0,
}))

function mergeRefs<T>(...refs: (Ref<T> | undefined | string)[]): RefCallback<T> {
    return (value) => {
        refs.forEach((ref) => {
            if (ref === undefined) {
                return
            }
            if (typeof ref === 'string') {
                return
            }
            if (typeof ref === 'function') {
                ref(value)
            } else if (ref !== null) {
                (ref as MutableRefObject<T | null>).current = value
            }
        })
    }
}

export interface CustomTextAreaProps {
    name?: string;

    ref?: UseFormRegisterReturn['ref'];
    onChange?: ChangeEventHandler;
    onBlur?: FocusEventHandler;

    onUserChange?: TextareaHTMLAttributes<HTMLTextAreaElement>['onChange'];
    onUserBlur?: TextareaHTMLAttributes<HTMLTextAreaElement>['onBlur'];

    isEmpty?: boolean;
    isFocused?: boolean;

    label: string
}

export default forwardRef<HTMLTextAreaElement, CustomTextAreaProps>(({
    // !**** Foward ref omits refs. But react hook form (RHF) inserts a ref, and we already have one 
    // So in the end we merge both in a function that calls them. That's why the error
    ref: rhfRef,         // The ref from RHF's register
    onChange: rhfOnChange, // The onChange from RHF's register
    onBlur: rhfOnBlur,     // The onBlur from RHF's register

    onUserChange,
    onUserBlur,

    label,
    ...props
}: CustomTextAreaProps, parentRef) => {
    const [isFocused, setIsFocused] = useState(false)

    const ref = useRef<HTMLTextAreaElement>(null)

    const refs = mergeRefs(ref, rhfRef, parentRef)

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        if (onUserChange) onUserChange(e);
        if (rhfOnChange) rhfOnChange(e);
    }

    const handleBlur = (e: React.FocusEvent<HTMLTextAreaElement>) => {
        setIsFocused(false);
        if (onUserBlur) onUserBlur(e);
        if (rhfOnBlur) rhfOnBlur(e);
    }

    return (
        <Box position="relative" width="100%">
            <Label
                isFocused={isFocused}
                isEmpty={Boolean(ref.current?.value)}
            >
                {label}
            </Label>
            <FieldSet
                height={ref.current?.clientHeight}
                isFocused={isFocused}
                isEmpty={Boolean(ref.current?.value)}
            >
                <TextArea
                    aria-placeholder="Insira a descrição da questão"
                    ref={refs}
                    rows={6}

                    onFocus={() => setIsFocused(true)}
                    onBlur={handleBlur}
                    onChange={handleChange}
                    isFocused={isFocused}
                    isEmpty={Boolean(ref.current?.value)}
                    {...props}
                />
                <legend style={{
                    position: "relative",
                    top: -50,
                    left: -10,

                    fontSize: 12,
                    height: 0,
                    color: "transparent",
                }}>
                    Insira o conteúdo da questão
                </legend>
            </FieldSet>
        </Box>
    )
}
)
// onSelect={(e) => {
//     e.target.style.border = "solid 2px " + theme.palette.primary.purpleDark
// }}
// onBlur={(e) => {
//     e.target.style.border = "solid 2px transparent"
// }}
// onMouseOver={(e) => {
//     e.target.style.border = "solid 1px rgb(180,180,180,0.57)"
// }}
// onMouseOut={(e) => {
//     e.target.style.border = "solid 1px rgb(180,180,180,0.27)"
// }}