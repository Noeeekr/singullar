// Components
import Stack, { StackProps } from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import CircularButton from "@components/buttons/Circular";

// Features
import { useNavigate } from "react-router-dom";
import TextFallback from "./TextFallback";

interface SectionHeaderProps extends StackProps {
  title: string;
  subtitle?: string;
  children?: JSX.Element;
  breadcrumbs?: JSX.Element;
}
/**
 * Contains the title, a button to navigate back and a text subtitle that is a breadcrumbs by default.
 *
 * Loads children in right side.
 */
const SectionHeader = ({ children, title, subtitle, breadcrumbs, ...props }: SectionHeaderProps): JSX.Element => {
  const navigate = useNavigate();

  return (
    <Stack direction="row" alignItems="center" width="100%" justifyContent="space-between" gap={1} {...props}>
      <Stack direction="row" alignItems="center" gap={1}>
        <CircularButton onClick={() => { navigate(-1) }} />
        <Stack gap={0.5} direction="column" alignItems="start">
          <TextFallback on={title == ""}>
            <Typography component="h3" variant="h4" fontWeight="600">
              {title}
            </Typography>
          </TextFallback>
          <Stack direction="row" gap={0.5}>
            {
              subtitle != undefined
                ? <TextFallback on={subtitle == ""} width={2}>
                  <Typography fontWeight={500} sx={{ color: 'rgb(130,130,130)' }}>
                    {subtitle}
                  </Typography>
                </TextFallback>
                : breadcrumbs
            }
          </Stack>
        </Stack>
      </Stack>
      {children || <></>}
    </Stack>
  );
};

export default SectionHeader;
