import { createSelector } from "reselect";
import { useAppSelector } from "../../store/store";
import { Note } from "../../models/NotesModels";
import { selectNotesList } from "./selectors";
import {
  selectUnsarchivedNotesList,
  selectArchivedNotesList,
} from "./selectors";

export const useNotes = (): Array<Note> =>
  useAppSelector(createSelector(selectNotesList, (i) => i));

export const useUnarchivedNotes = (): Array<Note> =>
  useAppSelector(createSelector(selectUnsarchivedNotesList, (i) => i));

export const useArchivedNotes = (): Array<Note> =>
  useAppSelector(createSelector(selectArchivedNotesList, (i) => i));
