export const ui = {
  pageHeader: 'mb-6 flex items-start justify-between gap-4 max-md:flex-col max-md:items-stretch',
  pageEyebrow: 'mb-1 text-xs font-semibold uppercase tracking-normal text-gray-500',
  pageTitle: 'm-0 text-[28px] font-semibold leading-tight text-gray-900 max-md:text-2xl',
  pageSubtitle: 'mt-2 max-w-2xl text-sm text-gray-500',
  card: 'rounded-2xl border border-gray-100 bg-white shadow-sm',
  sectionTitle: 'm-0 text-lg font-semibold text-gray-900',
  btnBase:
    'inline-flex min-h-10 items-center justify-center gap-2 rounded-xl border px-3.5 py-2 text-sm font-semibold leading-none transition-colors disabled:cursor-not-allowed disabled:opacity-60',
  btnPrimary: 'border-blue-600 bg-blue-600 text-white hover:bg-blue-700',
  btnSecondary: 'border-gray-200 bg-white text-gray-700 hover:border-gray-300 hover:bg-gray-50',
  btnDanger: 'border-red-100 bg-red-50 text-red-700 hover:border-red-200',
  btnSuccess: 'border-emerald-100 bg-emerald-50 text-emerald-700 hover:border-emerald-200',
  btnGhost: 'border-transparent bg-transparent text-gray-600 hover:bg-gray-100 hover:text-gray-900',
  btnIcon: 'h-10 w-10 min-w-10 px-0',
  formGrid: 'grid grid-cols-2 gap-4 max-sm:grid-cols-1',
  formField: 'flex flex-col gap-2',
  formFieldFull: 'col-span-full',
  formLabel: 'text-[13px] font-semibold text-gray-900',
  formControl:
    'w-full rounded-xl border border-gray-200 bg-white px-3 py-2.5 text-sm text-gray-700 outline-none transition focus:border-blue-400 focus:ring-2 focus:ring-blue-100',
  formTextarea:
    'min-h-24 w-full resize-y rounded-xl border border-gray-200 bg-white px-3 py-2.5 text-sm text-gray-700 outline-none transition focus:border-blue-400 focus:ring-2 focus:ring-blue-100',
  formHelp: 'm-0 text-xs text-gray-500',
  emptyState:
    'grid min-h-56 place-items-center rounded-2xl border border-dashed border-gray-200 bg-gray-50 p-7 text-center',
  emptyTitle: 'mb-2 text-lg font-semibold text-gray-900',
  emptyCopy: 'mx-auto max-w-md text-sm text-gray-600',
  tableWrap: 'overflow-x-auto rounded-2xl',
  table: 'w-full border-collapse text-sm',
  th: 'border-b border-gray-50 bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase text-gray-400',
  td: 'border-b border-gray-50 px-5 py-4 align-middle text-gray-900',
}

export const btn = (
  variant: keyof Pick<
    typeof ui,
    'btnPrimary' | 'btnSecondary' | 'btnDanger' | 'btnSuccess' | 'btnGhost'
  >,
) => `${ui.btnBase} ${ui[variant]}`
