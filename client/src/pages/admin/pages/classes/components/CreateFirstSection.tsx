// Components
import Stack from "@mui/material/Stack";

import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import ButtonSolid from '@components/ButtonSolid';

// Features
import { useContext, useEffect, useState } from "react";
import { CreateFormContext } from "./CreateForm";
import { parseNumIntoText } from './StudentTable';

// Types
interface IClassConfig {
  currentYear: number,
  segment: number,
  series: number,
}

/**
 * No students yet, no notifications, no allowed teachers..
 */
interface IClassNewborn extends IClassConfig {
  name: string,
  students: any[]
}

export interface IFirstSectionData extends IClassConfig {
  class: IClassNewborn
}

export const ButtonAddClass = (): JSX.Element => {
  const { formState, setFormValue, setSectionStatus } = useContext(CreateFormContext);
  const [disabled, setDisabled] = useState(true);

  useEffect(() => {
    if (formState.series != -1 && formState.segment != -1 && formState.currentYear != -1) {
      setDisabled(false)
    } else {
      setDisabled(true)
    }
  },[JSON.stringify(formState)]);

  return(
    <div style={{ margin: '0 0 0 auto'}}>
      <ButtonSolid 
        disabled={disabled}
        onClick={() => {
          setFormValue("class",{ 
              name: parseNumIntoText(1,formState.segment) + " " + parseNumIntoText(2, formState.series) + " " + formState.currentYear + " A",
              students: [],
              segment: formState.segment, 
              series: formState.series, 
              currentYear: formState.currentYear,
          });
          setFormValue("segment", -1)
          setFormValue("series", -1)
          setFormValue("currentYear", -1)
          setSectionStatus(2)
        }}
      >
        Adicionar turma
      </ButtonSolid>
    </div>
  )
}

const FirstSectionInputs = (): JSX.Element => {
  const { formState, setFormValue } = useContext(CreateFormContext);

  // Handle year value for first input
  const time = new Date().getFullYear();

  const years: number[] = [];
  for (let i = 2024; i <= time; i++) {
    years.push(i);
  }
    
  return (
    <Stack direction="column" gap={1}>
      <Stack direction="row" gap={2} padding={4}>
        <FormControl>
          <div style={{ position: "relative", width: "100%" }}>
            <InputLabel id="class-creation-year-input-label">
              Ano letivo
            </InputLabel>

            {/* React form hook controller */}
                <Select
                  labelId="class-creation-year-input-label"
                  label="Ano letivo"
                  aria-label="Ano letivo"
                  value={formState.currentYear}
                  onChange={(e) => (setFormValue("currentYear",Number(e.target.value)))}
                >
                  <MenuItem value={-1}>Escolha uma opção</MenuItem>
                  {years.map((year) => {
                    return <MenuItem key={year} value={year}>{year}</MenuItem>;
                  })}
                </Select>
          </div>
        </FormControl>
        <FormControl>
          <div style={{ position: "relative", width: "100%" }}>
            <InputLabel id="class-creation-segment-input-label">
              Segmento
            </InputLabel>
                  <Select
                      value={formState.segment}
                      onChange={(e) => (setFormValue("segment",Number(e.target.value)))}
                      labelId="class-creation-segment-input-label"
                      label="Segmento"
                      disabled={formState.currentYear === -1}
                      aria-label="Segmento"
                    >
                      <MenuItem value={-1}>Escolha uma opção</MenuItem>
                      <MenuItem value={1}>Ensino Fundamental 1</MenuItem>
                      <MenuItem value={2}>Ensino Fundamental 2</MenuItem>
                      <MenuItem value={3}>Ensino Médio</MenuItem>
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
                  {formState.segment === -1 ? (
                      <Select
                      value={formState.series}
                      onChange={(e) => (setFormValue("series",Number(e.target.value)))}
                          labelId="class-creation-series-input-label"
                        label="Série/Ano"
                        disabled={formState.segment === -1}
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
                        <MenuItem value={-1}>...</MenuItem>
                      </Select>
                    ) : (
                      <></>
                    )}
                    {formState.segment === 1 ? (
                      <Select
                      value={formState.series}
                      onChange={(e) => (setFormValue("series",Number(e.target.value)))}
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
                        <MenuItem value={0}>C.A (Classe de alfabetização)</MenuItem>
                        <MenuItem value={1}>1º Ano - Ensino Fundamental I</MenuItem>
                        <MenuItem value={2}>2º Ano - Ensino Fundamental I</MenuItem>
                        <MenuItem value={3}>3º Ano - Ensino Fundamental I</MenuItem>
                        <MenuItem value={4}>4º Ano - Ensino Fundamental I</MenuItem>
                        <MenuItem value={5}>5º Ano - Ensino Fundamental I</MenuItem>
                      </Select>
                    ) : (
                      <></>
                    )}
                    {formState.segment === 2 ? (
                      <Select
                      value={formState.series}
                      onChange={(e) => (setFormValue("series",Number(e.target.value)))}
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
                        <MenuItem value={6}>6º Ano - Ensino Fundamental II</MenuItem>
                        <MenuItem value={7}>7º Ano - Ensino Fundamental II</MenuItem>
                        <MenuItem value={8}>8º Ano - Ensino Fundamental II</MenuItem>
                        <MenuItem value={9}>9º Ano - Ensino Fundamental II</MenuItem>
                      </Select>
                    ) : (
                      <></>
                    )}
                    {formState.segment === 3 ? (
                      <Select
                      value={formState.series}
                      onChange={(e) => (setFormValue("series",Number(e.target.value)))}
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
                        <MenuItem value={10}>1º Ano - Ensino Médio</MenuItem>
                        <MenuItem value={11}>2º Ano - Ensino Médio</MenuItem>
                        <MenuItem value={12}>3º Ano - Ensino Médio</MenuItem>
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
