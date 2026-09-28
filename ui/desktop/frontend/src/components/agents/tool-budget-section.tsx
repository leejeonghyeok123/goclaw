import { useTranslation } from 'react-i18next'
import { numOrUndef } from '../../lib/format'
import type { ToolBudgetConfig } from '../../types/agent'

interface ToolBudgetSectionProps {
  value: ToolBudgetConfig
  onChange: (v: ToolBudgetConfig) => void
}

const FIELDS: { key: keyof ToolBudgetConfig; placeholder: string }[] = [
  { key: 'max_parallel_tool_calls', placeholder: '3' },
  { key: 'tool_result_max_tokens', placeholder: '1500' },
  { key: 'tool_loop_same_call_warning', placeholder: '3' },
  { key: 'tool_loop_same_call_critical', placeholder: '5' },
  { key: 'tool_loop_same_result_warning', placeholder: '4' },
  { key: 'tool_loop_same_result_critical', placeholder: '6' },
]

export function ToolBudgetSection({ value, onChange }: ToolBudgetSectionProps) {
  const { t } = useTranslation('agents')
  return (
    <div className="space-y-3">
      <div className="space-y-0.5">
        <h4 className="text-xs font-semibold text-text-primary">{t('configSections.toolBudget.title')}</h4>
        <p className="text-[11px] text-text-muted">{t('configSections.toolBudget.description')}</p>
      </div>
      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2">
        {FIELDS.map((field) => (
          <div key={field.key} className="space-y-1">
            <label className="text-[11px] font-medium text-text-secondary">
              {t(`configSections.toolBudget.${field.key}`)}
            </label>
            <input
              type="number"
              min={0}
              value={value[field.key] ?? ''}
              placeholder={field.placeholder}
              onChange={(e) => onChange({ ...value, [field.key]: numOrUndef(e.target.value) })}
              className="w-full bg-surface-tertiary border border-border rounded-lg px-3 py-2 text-base md:text-sm text-text-primary focus:outline-none focus:ring-1 focus:ring-accent"
            />
          </div>
        ))}
      </div>
    </div>
  )
}
