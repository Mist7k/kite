// Code generated by tygo. DO NOT EDIT.

//////////
// source: data.go

export interface MessageData {
  content?: string;
  flags?: number /* int */;
  attachments?: MessageAttachment[];
  embeds?: EmbedData[];
  components?: ComponentRowData[];
  allowed_mentions?: AllowedMentionsData;
}
export interface MessageAttachment {
  asset_id?: string;
}
export interface EmbedData {
  id?: number /* int */;
  title?: string;
  description?: string;
  url?: string;
  timestamp?: string /* RFC3339 */;
  color?: number /* int */;
  footer?: EmbedFooterData;
  image?: EmbedImageData;
  thumbnail?: EmbedThumbnailData;
  author?: EmbedAuthorData;
  fields?: EmbedFieldData[];
}
export interface EmbedFooterData {
  text?: string;
  icon_url?: string;
}
export interface EmbedImageData {
  url?: string;
}
export interface EmbedThumbnailData {
  url?: string;
}
export interface EmbedAuthorData {
  name?: string;
  url?: string;
  icon_url?: string;
}
export interface EmbedFieldData {
  id?: number /* int */;
  name?: string;
  value?: string;
  inline?: boolean;
}
export interface ComponentRowData {
  id?: number /* int */;
  components?: ComponentData[];
}
export interface ComponentData {
  id?: number /* int */;
  type?: number /* int */;
  disabled?: boolean;
  /**
   * Button
   */
  style?: number /* int */;
  label?: string;
  emoji?: ComponentEmojiData;
  url?: string;
  /**
   * Select Menu
   */
  placeholder?: string;
  min_values?: number /* int */;
  max_values?: number /* int */;
  options?: ComponentSelectOptionData[];
  flow_source_id?: string;
}
export interface ComponentSelectOptionData {
  id?: number /* int */;
  label?: string;
  description?: string;
  emoji?: ComponentEmojiData;
  default?: boolean;
  flow_source_id?: string;
}
export interface ComponentEmojiData {
  name?: string;
  id?: string;
  animated?: boolean;
}
export interface AllowedMentionsData {
  parse?: string[];
}
