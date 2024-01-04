export const selectNoteState = (state) => state?.notes ?? null;
export const selectNotesList = (state) =>
  selectNoteState(state)?.notesList ?? [];

export const selectUnsarchivedNotesList = (state) =>
  selectNotesList(state).filter(({ isArchived }) => !isArchived) ?? [];

export const selectArchivedNotesList = (state) =>
  selectNotesList(state).filter(({ isArchived }) => isArchived) ?? [];
