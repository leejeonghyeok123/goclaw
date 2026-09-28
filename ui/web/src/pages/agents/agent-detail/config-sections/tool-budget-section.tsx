import { useTranslation } from "react-i18next";
import { Input } from "@/components/ui/input";
import { InfoLabel } from "@/components/shared/info-label";
import type { ToolBudgetConfig } from "@/types/agent";
import { numOrUndef } from "./config-section";

interface ToolBudgetSectionProps {
  value: ToolBudgetConfig;
  onChange: (v: ToolBudgetConfig) => void;
}

const FIELDS: { key: keyof ToolBudgetConfig; placeholder: string }[] = [
  { key: "max_parallel_tool_calls", placeholder: "3" },
  { key: "tool_result_max_tokens", placeholder: "1500" },
  { key: "tool_loop_same_call_warning", placeholder: "3" },
  { key: "tool_loop_same_call_critical", placeholder: "5" },
  { key: "tool_loop_same_result_warning", placeholder: "4" },
  { key: "tool_loop_same_result_critical", placeholder: "6" },
];

export function ToolBudgetSection({ value, onChange }: ToolBudgetSectionProps) {
  const { t } = useTranslation("agents");
  return (
    <div className="space-y-3">
      <div className="space-y-0.5">
        <h4 className="text-sm font-medium">{t("configSections.toolBudget.title")}</h4>
        <p className="text-xs text-muted-foreground">{t("configSections.toolBudget.description")}</p>
      </div>
      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2">
        {FIELDS.map((field) => (
          <div key={field.key} className="grid gap-1.5">
            <InfoLabel tip={t(`configSections.toolBudget.${field.key}Tip`)}>
              {t(`configSections.toolBudget.${field.key}`)}
            </InfoLabel>
            <Input
              type="number"
              min={0}
              className="text-base md:text-sm"
              value={value[field.key] ?? ""}
              placeholder={field.placeholder}
              onChange={(e) => onChange({ ...value, [field.key]: numOrUndef(e.target.value) })}
            />
          </div>
        ))}
      </div>
    </div>
  );
}
