/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#include <future>
#include <memory>
#include <utility>
#include <stdio.h>

#include <glog/logging.h>
#include <gtest/gtest.h>

#include "SessiondMocks.h"

#include <google/protobuf/util/message_differencer.h>

using ::testing::Test;

namespace magma {

MATCHER_P(CheckCount, count, "") {
  int arg_count = arg.size();
  return arg_count == count;
}

MATCHER_P(CheckRuleNames, list_static_rules, "") {
  std::vector<PolicyRule> rules = arg;
  if (rules.size() != list_static_rules.size()) {
    return false;
  }
  for (PolicyRule rule : rules) {
    bool found = false;
    for (const std::string rule_to_check : list_static_rules) {
      if (rule.id() == rule_to_check) {
        found = true;
        break;
      }
    }
    if (!found) {
      return false;
    }
  }
  return true;
}

MATCHER_P(CheckTeids, configured_teids, "") {
  Teids pipelined_req_teids = static_cast<const Teids>(arg);

  if ((pipelined_req_teids.agw_teid() == configured_teids.agw_teid()) &&
      (pipelined_req_teids.enb_teid() == configured_teids.enb_teid())) {
    return true;
  }

  return false;
}

MATCHER_P2(CheckUpdateRequestCount, monitorCount, chargingCount, "") {
  auto req = static_cast<const UpdateSessionRequest>(arg);
  return req.updates().size() == chargingCount &&
         req.usage_monitors().size() == monitorCount;
}

MATCHER_P(CheckUpdateRequestNumber, request_number, "") {
  auto request = static_cast<const UpdateSessionRequest&>(arg);
  for (const auto& credit_usage_update : request.updates()) {
    int req_number = credit_usage_update.request_number();
    return req_number == request_number;
  }
  return false;
}

MATCHER_P(CheckCoreRequest, expected_request, "") {
  auto req    = static_cast<const CreateSessionRequest&>(arg);
  auto ex_req = static_cast<const CreateSessionRequest&>(expected_request);
  if (!google::protobuf::util::MessageDifferencer::Equals(
          ex_req.requested_units(), req.requested_units())) {
    return false;
  }

  // Add other check for the request
  return true;
}

MATCHER_P3(CheckTerminateRequestCount, imsi, monitorCount, chargingCount, "") {
  auto req = static_cast<const SessionTerminateRequest>(arg);
  return req.common_context().sid().id() == imsi &&
         req.credit_usages().size() == chargingCount &&
         req.monitor_usages().size() == monitorCount;
}

MATCHER_P6(
    CheckSessionInfos, imsi_list, ip_address_list, ipv6_address_list, cfg,
    static_rule_lists, dynamic_rule_ids_lists, "") {
  auto infos = static_cast<const std::vector<SessionState::SessionInfo>>(arg);

  if (infos.size() != imsi_list.size()) return false;

  for (size_t i = 0; i < infos.size(); i++) {
    if (infos[i].imsi != imsi_list[i]) return false;
    if (infos[i].ip_addr != ip_address_list[i]) return false;
    if (infos[i].ipv6_addr != ipv6_address_list[i]) return false;
    if (infos[i].static_rules.size() != static_rule_lists[i].size())
      return false;
    if (infos[i].gx_dynamic_rules.size() != dynamic_rule_ids_lists[i].size())
      return false;
    for (size_t r_index = 0; i < infos[i].static_rules.size(); i++) {
      if (infos[i].static_rules[r_index] != static_rule_lists[i][r_index])
        return false;
    }
    for (size_t r_index = 0; i < infos[i].gx_dynamic_rules.size(); i++) {
      if (infos[i].gx_dynamic_rules[r_index].id() !=
          dynamic_rule_ids_lists[i][r_index])
        return false;
    }
    // check ambr field if config has qos_info
    if (cfg.rat_specific_context.has_lte_context() &&
        cfg.rat_specific_context.lte_context().has_qos_info()) {
      const auto& qos_info = cfg.rat_specific_context.lte_context().qos_info();
      if (!infos[i].ambr) {
        return false;
      } else if (infos[i].ambr->max_bandwidth_ul() != qos_info.apn_ambr_ul()) {
        return false;
      } else if (infos[i].ambr->max_bandwidth_dl() != qos_info.apn_ambr_dl()) {
        return false;
      }
    }
  }
  return true;
}

MATCHER_P(CheckEventType, expectedEventType, "") {
  return (arg.event_type() == expectedEventType);
}

MATCHER_P2(CheckCreateBearerReq, imsi, rule_count, "") {
  auto request = static_cast<const CreateBearerRequest>(arg);
  return request.sid().id() == imsi &&
         request.policy_rules().size() == rule_count;
}

MATCHER_P3(CheckDeleteOneBearerReq, imsi, link_bearer_id, eps_bearer_id, "") {
  auto request = static_cast<const DeleteBearerRequest>(arg);

  return request.sid().id() == imsi &&
         request.link_bearer_id() == uint32_t(link_bearer_id) &&
         request.eps_bearer_ids_size() == 1 &&
         request.eps_bearer_ids(0) == uint32_t(eps_bearer_id);
}

MATCHER_P(CheckSubset, ids, "") {
  auto request = static_cast<const std::vector<PolicyRule>>(arg);
  for (size_t i = 0; i < request.size(); i++) {
    if (ids.find(request[i].id()) != ids.end()) {
      return true;
    }
  }
  return false;
}

MATCHER_P(CheckPolicyID, id, "") {
  auto request = static_cast<const std::vector<PolicyRule>>(arg);
  for (size_t i = 0; i < request.size(); i++) {
    if (request[i].id() == id) {
      return true;
    }
  }
  return false;
}

MATCHER_P2(CheckPolicyIDs, count, ids, "") {
  auto request = static_cast<const std::vector<PolicyRule>>(arg);
  if (request.size() != (unsigned int) count) {
    return false;
  }
  for (size_t i = 0; i < request.size(); i++) {
    if (ids.find(request[i].id()) != ids.end()) {
      return true;
    }
  }
  return false;
}

MATCHER_P(CheckSubscriberQuotaUpdate, quota, "") {
  auto update = static_cast<std::vector<SubscriberQuotaUpdate>>(arg);
  if (update.size() != 1) {
    return false;
  }
  std::cerr << "\n\n" << update[0].update_type() << " \n\n";
  return update[0].update_type() == quota;
}

MATCHER_P2(CheckCreateSession, imsi, promise_p, "") {
  auto req = static_cast<const CreateSessionRequest*>(arg);
  promise_p->set_value(req->session_id());
  auto res = req->common_context().sid().id() == imsi;
  return res;
}

MATCHER_P(CheckSingleUpdate, expected_update, "") {
  auto request = static_cast<const UpdateSessionRequest*>(arg);
  if (request->updates_size() != 1) {
    return false;
  }

  auto& update = request->updates(0);
  bool val =
      update.usage().type() == expected_update.usage().type() &&
      update.usage().bytes_tx() == expected_update.usage().bytes_tx() &&
      update.usage().bytes_rx() == expected_update.usage().bytes_rx() &&
      update.common_context().sid().id() ==
          expected_update.common_context().sid().id() &&
      update.usage().charging_key() == expected_update.usage().charging_key();
  return val;
}

MATCHER_P(CheckTerminate, imsi, "") {
  auto request = static_cast<const SessionTerminateRequest*>(arg);
  return request->common_context().sid().id() == imsi;
}

MATCHER_P6(
    CheckActivateFlowsForTunnIds, imsi, ipv4, ipv6, enb_teid, agw_teid,
    rule_count, "") {
  auto request = static_cast<const ActivateFlowsRequest*>(arg);
  std::cerr << "Got dynamic size : " << request->dynamic_rules_size()
            << std::endl;
  auto res = request->sid().id() == imsi && request->ip_addr() == ipv4 &&
             request->ipv6_addr() == ipv6 &&
             request->uplink_tunnel() == agw_teid &&
             request->downlink_tunnel() == enb_teid &&
             request->dynamic_rules_size() == rule_count;
  return res;
}

MATCHER_P(CheckDeactivateFlows, imsi, "") {
  auto request = static_cast<const DeactivateFlowsRequest*>(arg);
  return request->sid().id() == imsi;
}

};  // namespace magma