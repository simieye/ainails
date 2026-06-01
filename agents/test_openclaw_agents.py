import unittest
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parent))

from openclaw_agents import MATRIX, execute_workflow


class AgentMatrixTest(unittest.TestCase):
    def test_matrix_contains_sixty_four_agents(self):
        self.assertEqual(len(MATRIX), 64)
        self.assertEqual(MATRIX[0].id, "agent-01-01")

    def test_execute_growth_workflow_returns_steps_and_roi(self):
        result = execute_workflow(
            {
                "goal": "测试美国市场增长飞轮",
                "workflow_id": "wf-growth-flywheel",
                "campaign_id": "cmp-test",
                "context": {
                    "market": "United States",
                    "product": "AI phone",
                    "budget_usd": 2000,
                },
            }
        )
        self.assertEqual(result["campaign_id"], "cmp-test")
        self.assertGreaterEqual(len(result["steps"]), 6)
        self.assertGreater(result["roi"], 0)
        self.assertIn("OpenClaw 已调度", result["summary"])


if __name__ == "__main__":
    unittest.main()
