import Phases from "./Phases.vue";
import i18n from "../i18n";
import "../tooltip";

export default {
  title: "Main/Phases",
  component: Phases,
  argTypes: {},
};

const Template = (args, { argTypes }) => ({
  i18n,
  props: Object.keys(argTypes),
  components: { Phases },
  template: '<Phases v-bind="$props"></Phases>',
});

export const Base = Template.bind({});
Base.args = {
  activePhases: 1,
  minCurrent: 6,
  maxCurrent: 32,
  chargeCurrent: 10,
};

export const OnePhase = Template.bind({});
OnePhase.args = {
  activePhases: 1,
  minCurrent: 6,
  maxCurrent: 16,
  chargeCurrent: 6,
};

export const TwoPhase = Template.bind({});
TwoPhase.args = {
  activePhases: 2,
  minCurrent: 6,
  maxCurrent: 16,
  chargeCurrent: 8,
};

export const ThreePhases = Template.bind({});
ThreePhases.args = {
  activePhases: 3,
  minCurrent: 6,
  maxCurrent: 16,
  chargeCurrent: 12,
};

export const RealCurrents = Template.bind({});
RealCurrents.args = {
  activePhases: 3,
  minCurrent: 6,
  maxCurrent: 32,
  chargeCurrent: 32,
  chargeCurrents: [15.2, 15.76, 14.5],
};
