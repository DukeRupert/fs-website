export interface CtaData {
    title: string;
    description: string;
    primary_action: Action;
    secondary_action: Action | null;
}

export interface Action {
    name: string;
    href: string;
}