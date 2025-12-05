// Components
import Stack from "@mui/material/Stack";

import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import ButtonSolid from '@components/buttons/Default/Solid';

// Features
import { useContext, useEffect, useState, useCallback } from "react";
import { FormContext } from "./Form";

// Types
import type { UserSegments } from "@models/server";
import { EF1, EF2, EM } from "@models/server";

export type FormFirstSectionData = {
  className: string,
  series: string,
  segment: UserSegments,
}

export const ButtonAddClass = (): JSX.Element => {
  const { formSections, setFormData, setActiveSection } = useContext(FormContext);
  const [isDisabled, setIsDisabled] = useState(Boolean(formSections.firstSection.segment));

  useEffect(() => {
    if (formSections.firstSection.segment) {
      setIsDisabled(false)
    } else {
      setIsDisabled(true)
    }
  }, [formSections]);

  const enableNextSection = useCallback(() => {
    if (!formSections.firstSection.segment) {
      return
    }
    setFormData("firstSection.className", `${formSections.firstSection.segment} ${formSections.firstSection.series}`);
    setActiveSection(2)
  }, [formSections, setFormData, setActiveSection])

  return (
    <div style={{ margin: '0 0 0 auto' }}>
      <ButtonSolid
        title="Adicionar Turma"
        disabled={isDisabled}
        onClick={enableNextSection}
      />
    </div>
  )
}

const FirstSectionInputs = (): JSX.Element => {
  const { formSections, setFormData } = useContext(FormContext);

  return (
    <Stack direction="column" gap={1}>
      <Stack direction="row" gap={2} padding={4}>
        <FormControl>
          <div style={{ position: "relative", width: "100%" }}>
            <InputLabel id="class-creation-segment-input-label">
              Segmento
            </InputLabel>
            <Select
              value={formSections.firstSection.segment}
              onChange={(e) => { if (e.target.value != null) setFormData("firstSection.segment", e.target.value as UserSegments) }}
              labelId="class-creation-segment-input-label"
              label="Segmento"
              aria-label="Segmento"
            >
              <MenuItem value={""}>Escolha uma opção</MenuItem>
              <MenuItem value={EF1}>Ensino Fundamental 1</MenuItem>
              <MenuItem value={EF2}>Ensino Fundamental 2</MenuItem>
              <MenuItem value={EM}>Ensino Médio</MenuItem>
            </Select>
          </div>
        </FormControl>
        <FormControl>
          <div style={{ position: "relative", width: "100%" }}>
            <InputLabel id="class-creation-series-input-label">
              Série/Ano
            </InputLabel>
            <>
              {/*
                      Different Select forms for each options of the previous input. That way
                      the user won't be able to select an invalid option. 
                  */}
              {formSections.firstSection.segment == "" || formSections.firstSection.segment == undefined ? (
                <Select
                  value={formSections.firstSection.series}
                  onChange={(e) => (setFormData("firstSection.series", e.target.value))}
                  labelId="class-creation-series-input-label"
                  label="Série/Ano"
                  disabled={true}
                  aria-label="Série/Ano"
                  MenuProps={{
                    PaperProps: {
                      sx: {
                        maxHeight: "200px",
                        marginTop: "10px",
                      },
                    },
                  }}
                >
                  <MenuItem value={""}>...</MenuItem>
                </Select>
              ) : (
                <></>
              )}
              {formSections.firstSection.segment == EF1 ? (
                <Select
                  value={formSections.firstSection.series}
                  onChange={(e) => (setFormData("firstSection.series", e.target.value))}
                  labelId="class-creation-series-input-label"
                  label="Série/Ano"
                  aria-label="Série/Ano"
                  MenuProps={{
                    PaperProps: {
                      sx: {
                        maxHeight: "200px",
                        marginTop: "10px",
                      },
                    },
                  }}
                >
                  <MenuItem value={"alfabetização"}>C.A (Classe de alfabetização)</MenuItem>
                  <MenuItem value={"1º Ano"}>1º Ano - Ensino Fundamental I</MenuItem>
                  <MenuItem value={"2º Ano"}>2º Ano - Ensino Fundamental I</MenuItem>
                  <MenuItem value={"3º Ano"}>3º Ano - Ensino Fundamental I</MenuItem>
                  <MenuItem value={"4º Ano"}>4º Ano - Ensino Fundamental I</MenuItem>
                  <MenuItem value={"5º Ano"}>5º Ano - Ensino Fundamental I</MenuItem>
                </Select>
              ) : (
                <></>
              )}
              {formSections.firstSection.segment == EF2 ? (
                <Select
                  value={formSections.firstSection.series}
                  onChange={(e) => (setFormData("firstSection.series", e.target.value))}
                  labelId="class-creation-series-input-label"
                  label="Série/Ano"
                  aria-label="Série/Ano"
                  MenuProps={{
                    PaperProps: {
                      sx: {
                        maxHeight: "200px",
                        marginTop: "10px",
                      },
                    },
                  }}
                >
                  <MenuItem value={"6º Ano"}>6º Ano - Ensino Fundamental II</MenuItem>
                  <MenuItem value={"7º Ano"}>7º Ano - Ensino Fundamental II</MenuItem>
                  <MenuItem value={"8º Ano"}>8º Ano - Ensino Fundamental II</MenuItem>
                  <MenuItem value={"9º Ano"}>9º Ano - Ensino Fundamental II</MenuItem>
                </Select>
              ) : (
                <></>
              )}
              {formSections.firstSection.segment == EM ? (
                <Select
                  value={formSections.firstSection.series}
                  onChange={(e) => (setFormData("firstSection.series", e.target.value))}
                  labelId="class-creation-series-input-label"
                  label="Série/Ano"
                  aria-label="Série/Ano"
                  MenuProps={{
                    PaperProps: {
                      sx: {
                        maxHeight: "200px",
                        marginTop: "10px",
                      },
                    },
                  }}
                >
                  <MenuItem value={"1º Ano"}>1º Ano - Ensino Médio</MenuItem>
                  <MenuItem value={"2º Ano"}>2º Ano - Ensino Médio</MenuItem>
                  <MenuItem value={"3º Ano"}>3º Ano - Ensino Médio</MenuItem>
                </Select>
              ) : (
                <></>
              )}
            </>
          </div>
        </FormControl>
      </Stack>
    </Stack>
  );
};

export default FirstSectionInputs;
