// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.
import {
  TransactionInsightDetails,
  TransactionInsightDetailsStateProps,
  TransactionInsightDetailsDispatchProps,
} from "./transactionInsightDetails";
import { connect } from "react-redux";
import { RouteComponentProps, withRouter } from "react-router-dom";
import { AppState, uiConfigActions } from "src/store";
import {
  selectTransactionInsightDetails,
  selectTransactionInsightDetailsError,
  actions,
  selectTransactionInsightDetailsMaxSizeReached,
} from "src/store/insightDetails/transactionInsightDetails";
import { TimeScale } from "../../timeScaleDropdown";
import { actions as sqlStatsActions } from "../../store/sqlStats";
import { Dispatch } from "redux";
import { TransactionInsightEventDetailsRequest } from "src/api";
import { selectHasAdminRole } from "src/store/uiConfig";
import { actions as analyticsActions } from "../../store/analytics";

const mapStateToProps = (
  state: AppState,
  props: RouteComponentProps,
): TransactionInsightDetailsStateProps => {
  const insightDetails = selectTransactionInsightDetails(state, props);
  const insightError = selectTransactionInsightDetailsError(state, props);
  return {
    insightEventDetails: insightDetails,
    insightError: insightError,
    hasAdminRole: selectHasAdminRole(state),
    maxSizeApiReached: selectTransactionInsightDetailsMaxSizeReached(
      state,
      props,
    ),
  };
};

const mapDispatchToProps = (
  dispatch: Dispatch,
): TransactionInsightDetailsDispatchProps => ({
  refreshTransactionInsightDetails: (
    req: TransactionInsightEventDetailsRequest,
  ) => {
    dispatch(actions.refresh(req));
  },
  setTimeScale: (ts: TimeScale) => {
    dispatch(
      sqlStatsActions.updateTimeScale({
        ts: ts,
      }),
    );
    dispatch(
      analyticsActions.track({
        name: "TimeScale changed",
        page: "Transaction Insight Details",
        value: ts.key,
      }),
    );
  },
  refreshUserSQLRoles: () => dispatch(uiConfigActions.refreshUserSQLRoles()),
});

export const TransactionInsightDetailsConnected = withRouter(
  connect<
    TransactionInsightDetailsStateProps,
    TransactionInsightDetailsDispatchProps,
    RouteComponentProps
  >(
    mapStateToProps,
    mapDispatchToProps,
  )(TransactionInsightDetails),
);
