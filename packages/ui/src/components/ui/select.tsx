import {
  useCallback,
  useEffect,
  useId,
  useRef,
  useState,
  type KeyboardEvent,
  type ReactNode,
} from "react";

import { cn } from "../../lib";

export type SelectOption = {
  value: string;
  label: ReactNode;
  disabled?: boolean;
};

type SelectProps = {
  value?: string;
  defaultValue?: string;
  onChange?: (value: string) => void;
  options: SelectOption[];
  placeholder?: string;
  disabled?: boolean;
  className?: string;
  id?: string;
  name?: string;
  "aria-label"?: string;
  "aria-labelledby"?: string;
  "data-testid"?: string;
};

export function Select({
  value: controlledValue,
  defaultValue,
  onChange,
  options,
  placeholder = "Select...",
  disabled,
  className,
  id,
  name,
  "aria-label": ariaLabel,
  "aria-labelledby": ariaLabelledBy,
  "data-testid": dataTestId,
}: SelectProps) {
  const generatedId = useId();
  const componentId = id ?? generatedId;
  const listboxId = `${componentId}-listbox`;

  const isControlled = controlledValue !== undefined;
  const [internalValue, setInternalValue] = useState<string>(
    defaultValue ?? "",
  );
  const value = isControlled ? controlledValue : internalValue;

  const [open, setOpen] = useState(false);
  const [activeIndex, setActiveIndex] = useState<number>(() =>
    Math.max(
      0,
      options.findIndex((o) => o.value === (controlledValue ?? defaultValue)),
    ),
  );

  const rootRef = useRef<HTMLDivElement>(null);
  const triggerRef = useRef<HTMLButtonElement>(null);

  const selected = options.find((o) => o.value === value);

  const commit = useCallback(
    (next: string) => {
      if (!isControlled) setInternalValue(next);
      onChange?.(next);
    },
    [isControlled, onChange],
  );

  const close = useCallback(() => {
    setOpen(false);
    triggerRef.current?.focus();
  }, []);

  useEffect(() => {
    if (!open) return;
    const onDocClick = (event: MouseEvent) => {
      if (!rootRef.current?.contains(event.target as Node)) setOpen(false);
    };
    document.addEventListener("mousedown", onDocClick);
    return () => document.removeEventListener("mousedown", onDocClick);
  }, [open]);

  const moveActive = useCallback(
    (delta: number) => {
      if (options.length === 0) return;
      setActiveIndex((prev) => {
        let next = prev;
        for (let i = 0; i < options.length; i += 1) {
          next = (next + delta + options.length) % options.length;
          if (!options[next]?.disabled) return next;
        }
        return prev;
      });
    },
    [options],
  );

  const handleTriggerKeyDown = (event: KeyboardEvent<HTMLButtonElement>) => {
    if (disabled) return;
    if (
      event.key === "ArrowDown" ||
      event.key === "ArrowUp" ||
      event.key === "Enter" ||
      event.key === " "
    ) {
      event.preventDefault();
      setOpen(true);
    }
  };

  const handleListKeyDown = (event: KeyboardEvent<HTMLDivElement>) => {
    if (event.key === "Escape") {
      event.preventDefault();
      close();
    } else if (event.key === "ArrowDown") {
      event.preventDefault();
      moveActive(1);
    } else if (event.key === "ArrowUp") {
      event.preventDefault();
      moveActive(-1);
    } else if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      const opt = options[activeIndex];
      if (opt && !opt.disabled) {
        commit(opt.value);
        close();
      }
    } else if (event.key === "Tab") {
      setOpen(false);
    }
  };

  return (
    <div ref={rootRef} className={cn("ui-select-root", className)}>
      {name ? <input type="hidden" name={name} value={value} /> : null}
      <button
        ref={triggerRef}
        id={componentId}
        type="button"
        role="combobox"
        aria-haspopup="listbox"
        aria-expanded={open}
        aria-controls={listboxId}
        aria-disabled={disabled || undefined}
        aria-label={ariaLabel}
        aria-labelledby={ariaLabelledBy}
        disabled={disabled}
        data-testid={dataTestId}
        className={cn("ui-select", open && "ui-select--open")}
        onClick={() => !disabled && setOpen((v) => !v)}
        onKeyDown={handleTriggerKeyDown}
      >
        <span className="ui-select-value">
          {selected ? selected.label : placeholder}
        </span>
        <span aria-hidden="true" className="ui-select-arrow">
          {open ? "▲" : "▼"}
        </span>
      </button>
      {open && !disabled ? (
        <div
          id={listboxId}
          role="listbox"
          tabIndex={-1}
          aria-activedescendant={
            options[activeIndex]
              ? `${componentId}-option-${activeIndex}`
              : undefined
          }
          className="ui-select-menu"
          onKeyDown={handleListKeyDown}
          ref={(node) => {
            node?.focus();
          }}
        >
          {options.map((option, index) => {
            const isSelected = option.value === value;
            const isActive = index === activeIndex;
            return (
              <div
                key={option.value}
                id={`${componentId}-option-${index}`}
                role="option"
                aria-selected={isSelected}
                aria-disabled={option.disabled || undefined}
                className={cn(
                  "ui-select-option",
                  isSelected && "ui-select-option--selected",
                  isActive && "ui-select-option--active",
                  option.disabled && "ui-select-option--disabled",
                )}
                onMouseEnter={() => setActiveIndex(index)}
                onClick={() => {
                  if (option.disabled) return;
                  commit(option.value);
                  close();
                }}
              >
                {option.label}
              </div>
            );
          })}
        </div>
      ) : null}
    </div>
  );
}
